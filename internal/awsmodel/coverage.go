package awsmodel

import (
	"fmt"
	"sort"
	"strings"
)

// DispatchInventory is the gateway-side operation registration state supplied
// to CompareOperations.
type DispatchInventory struct {
	Registered  []string
	Stubbed     []string
	Unsupported []string
	Opaque      bool
	Note        string
}

// OperationCoverage is a deterministic model-versus-dispatch comparison.
// Implemented contains only modelled operations with a real handler; Extra
// contains registered operations absent from the pinned model.
type OperationCoverage struct {
	Service     Service
	APIVersion  string
	Modelled    []string
	Registered  []string
	Implemented []string
	Stubbed     []string
	Unsupported []string
	Missing     []string
	Extra       []string
	Opaque      bool
	Note        string
}

// CompareOperations compares one embedded service model with an authoritative
// gateway dispatch inventory.
func CompareOperations(service Service, dispatch DispatchInventory) (OperationCoverage, error) {
	model, err := Load(service)
	if err != nil {
		return OperationCoverage{}, err
	}
	coverage := OperationCoverage{
		Service:    service,
		APIVersion: model.Metadata().APIVersion,
		Modelled:   model.Operations(),
		Opaque:     dispatch.Opaque,
		Note:       dispatch.Note,
	}
	if dispatch.Opaque {
		return coverage, nil
	}

	registered, err := uniqueSorted("registered", dispatch.Registered)
	if err != nil {
		return OperationCoverage{}, fmt.Errorf("awsmodel: %s dispatch: %w", service, err)
	}
	stubbed, err := uniqueSorted("stubbed", dispatch.Stubbed)
	if err != nil {
		return OperationCoverage{}, fmt.Errorf("awsmodel: %s dispatch: %w", service, err)
	}
	unsupported, err := uniqueSorted("unsupported", dispatch.Unsupported)
	if err != nil {
		return OperationCoverage{}, fmt.Errorf("awsmodel: %s dispatch: %w", service, err)
	}
	registeredSet := toSet(registered)
	stubbedSet := toSet(stubbed)
	unsupportedSet := toSet(unsupported)
	for operation := range stubbedSet {
		if !registeredSet[operation] {
			return OperationCoverage{}, fmt.Errorf("awsmodel: %s stubbed operation %q is not registered", service, operation)
		}
		if unsupportedSet[operation] {
			return OperationCoverage{}, fmt.Errorf("awsmodel: %s operation %q is both stubbed and unsupported", service, operation)
		}
	}
	for operation := range unsupportedSet {
		if !registeredSet[operation] {
			return OperationCoverage{}, fmt.Errorf("awsmodel: %s unsupported operation %q is not registered", service, operation)
		}
	}

	modelledSet := toSet(coverage.Modelled)
	for _, operation := range coverage.Modelled {
		if !registeredSet[operation] {
			coverage.Missing = append(coverage.Missing, operation)
			continue
		}
		if !stubbedSet[operation] && !unsupportedSet[operation] {
			coverage.Implemented = append(coverage.Implemented, operation)
		}
	}
	for _, operation := range registered {
		if !modelledSet[operation] {
			coverage.Extra = append(coverage.Extra, operation)
		}
	}
	coverage.Registered = registered
	coverage.Stubbed = stubbed
	coverage.Unsupported = unsupported
	return coverage, nil
}

func uniqueSorted(label string, values []string) ([]string, error) {
	result := append([]string(nil), values...)
	sort.Strings(result)
	for index, value := range result {
		if value == "" {
			return nil, fmt.Errorf("%s operation is empty", label)
		}
		if index > 0 && result[index-1] == value {
			return nil, fmt.Errorf("%s operation %q is duplicated", label, value)
		}
	}
	return result, nil
}

func toSet(values []string) map[string]bool {
	set := make(map[string]bool, len(values))
	for _, value := range values {
		set[value] = true
	}
	return set
}

// RenderCoverageSummary renders the cross-service counts as plain text for a
// terminal. The published pages carry the per-operation detail.
func RenderCoverageSummary(coverages []OperationCoverage) string {
	var report strings.Builder
	fmt.Fprintf(&report, "AWS model operation coverage (aws-sdk-go %s)\n\n", SourceSDKVersion)
	for _, coverage := range sortedCoverages(coverages) {
		if coverage.Opaque {
			fmt.Fprintf(&report, "%-24s %-12s not enumerable\n", coverage.Service, coverage.APIVersion)
			continue
		}
		fmt.Fprintf(&report, "%-24s %-12s %3d of %3d implemented (%5.1f%%), %3d stubbed, %3d unsupported, %3d outside the model\n",
			coverage.Service, coverage.APIVersion, len(coverage.Implemented), len(coverage.Modelled),
			coverage.ImplementedPercent(), len(coverage.Stubbed), len(coverage.Unsupported), len(coverage.Extra))
	}
	return report.String()
}

// RenderInternalReport renders the full comparison, including the operations
// the published pages leave out. It is not a site page: the docs build renders
// only the slugs it is configured with, and this file is not one of them.
//
// A declared not-applicable operation is listed apart from the gaps, with the
// reason, so a list of candidate work never carries one nobody intends to
// build.
func RenderInternalReport(coverages []OperationCoverage, pages PageSet) string {
	var report strings.Builder
	report.WriteString("# AWS Model Operation Coverage (internal)\n\n")
	report.WriteString("Generated by `make aws-model-coverage`. This file is not published: it carries the counts and the unimplemented operations that the pages under this directory deliberately leave out.\n\n")
	report.WriteString("**Not implemented** is the candidate work. **Not applicable** is declared in `pages.json` and is not work: those operations describe something this platform will not have.\n\n")
	fmt.Fprintf(&report, "Pinned model source: `aws-sdk-go %s`.\n\n", SourceSDKVersion)

	for _, coverage := range sortedCoverages(coverages) {
		page := pages.Services[coverage.Service]
		fmt.Fprintf(&report, "## %s\n\n", coverage.Service)
		if coverage.Opaque {
			fmt.Fprintf(&report, "`%s` — not enumerable. %s\n\n", coverage.APIVersion, coverage.Note)
			continue
		}

		missing, notApplicable := splitNotApplicable(coverage.Missing, page)
		refused, declaredRefusals := splitNotApplicable(coverage.Unsupported, page)
		notApplicable = append(notApplicable, declaredRefusals...)
		sort.Strings(notApplicable)

		fmt.Fprintf(&report, "`%s` — %d of %d modelled operations implemented (%.1f%%), %d stubbed, %d not applicable, %d registered outside the pinned model.\n\n",
			coverage.APIVersion, len(coverage.Implemented), len(coverage.Modelled), coverage.ImplementedPercent(),
			len(coverage.Stubbed), len(notApplicable), len(coverage.Extra))
		writeOperationList(&report, "Not implemented", missing)
		writeOperationList(&report, "Stubbed", coverage.Stubbed)
		writeOperationList(&report, "Registered and refused", refused)
		writeOperationList(&report, "Registered outside the pinned model", coverage.Extra)
		writeNotApplicableList(&report, notApplicable, page)
	}
	return report.String()
}

// splitNotApplicable divides operations into those still open and those the
// page declares the platform will never serve.
func splitNotApplicable(operations []string, page PageMetadata) (open, notApplicable []string) {
	for _, operation := range operations {
		if _, declared := page.NotApplicable[operation]; declared {
			notApplicable = append(notApplicable, operation)
			continue
		}
		open = append(open, operation)
	}
	return open, notApplicable
}

// writeNotApplicableList records why each declared operation is not work,
// grouped by the reason it was declared under.
func writeNotApplicableList(report *strings.Builder, operations []string, page PageMetadata) {
	if len(operations) == 0 {
		return
	}
	byNote := map[string][]string{}
	for _, operation := range operations {
		key := page.NotApplicable[operation]
		byNote[key] = append(byNote[key], operation)
	}

	report.WriteString("### Not applicable — do not build\n\n")
	for _, key := range noteOrder(page.Notes) {
		if len(byNote[key]) == 0 {
			continue
		}
		fmt.Fprintf(report, "%s: %s\n\n", key, page.Notes[key])
		for _, operation := range byNote[key] {
			fmt.Fprintf(report, "- `%s`\n", operation)
		}
		report.WriteString("\n")
	}
}

func writeOperationList(report *strings.Builder, heading string, operations []string) {
	if len(operations) == 0 {
		return
	}
	fmt.Fprintf(report, "### %s\n\n", heading)
	for _, operation := range operations {
		fmt.Fprintf(report, "- `%s`\n", operation)
	}
	report.WriteString("\n")
}

func sortedCoverages(coverages []OperationCoverage) []OperationCoverage {
	ordered := append([]OperationCoverage(nil), coverages...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Service < ordered[j].Service })
	return ordered
}
