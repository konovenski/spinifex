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

func sortedCoverages(coverages []OperationCoverage) []OperationCoverage {
	ordered := append([]OperationCoverage(nil), coverages...)
	sort.Slice(ordered, func(i, j int) bool { return ordered[i].Service < ordered[j].Service })
	return ordered
}
