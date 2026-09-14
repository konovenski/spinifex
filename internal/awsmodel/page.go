package awsmodel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Published operation states. The symbol carries the status at a glance and the
// label disambiguates it; the docs site escapes inline SVG, so these are runes.
const (
	StatusImplemented    = "✅ Implemented"
	StatusStub           = "🟡 Stub"
	StatusNotSupported   = "🚫 Not supported"
	StatusNotImplemented = "❌ Not implemented"
	StatusOutsideModel   = "🔒 Outside the pinned model"
)

// Frontmatter length rules enforced by the docs site. See docs/README.md.
const (
	seoTitleSuffix = " — Spinifex Docs"
	seoTitleMin    = 50
	seoTitleMax    = 60
	descriptionMin = 150
	descriptionMax = 160
)

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

// OperationStatus is one row of a published coverage table.
type OperationStatus struct {
	Operation string
	Status    string
}

// ImplementedPercent is the share of modelled operations bound to a real
// handler, or zero for a model with no operations.
func (c OperationCoverage) ImplementedPercent() float64 {
	if len(c.Modelled) == 0 {
		return 0
	}
	return 100 * float64(len(c.Implemented)) / float64(len(c.Modelled))
}

// OperationStatuses returns every modelled operation with its dispatch state,
// followed by the registered operations the pinned model does not describe.
func (c OperationCoverage) OperationStatuses() []OperationStatus {
	registered := toSet(c.Registered)
	stubbed := toSet(c.Stubbed)
	unsupported := toSet(c.Unsupported)

	statuses := make([]OperationStatus, 0, len(c.Modelled)+len(c.Extra))
	for _, operation := range c.Modelled {
		status := StatusNotImplemented
		switch {
		case stubbed[operation]:
			status = StatusStub
		case unsupported[operation]:
			status = StatusNotSupported
		case registered[operation]:
			status = StatusImplemented
		}
		statuses = append(statuses, OperationStatus{Operation: operation, Status: status})
	}
	for _, operation := range c.Extra {
		statuses = append(statuses, OperationStatus{Operation: operation, Status: StatusOutsideModel})
	}
	return statuses
}

// PageMetadata is the checked-in frontmatter for one published coverage page.
// It is hand-maintained rather than derived, because a description carrying a
// live count would drift outside the length the docs site enforces.
type PageMetadata struct {
	Slug        string   `json:"slug"`
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	SEOTitle    string   `json:"seoTitle"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// PageSet is the metadata for the index page and every service page.
type PageSet struct {
	Category string                   `json:"category"`
	Index    PageMetadata             `json:"index"`
	Services map[Service]PageMetadata `json:"services"`
}

// ParsePageSet decodes docs/compatibility/pages.json and checks it describes
// exactly the services the loader knows about, with publishable frontmatter.
func ParsePageSet(contents []byte) (PageSet, error) {
	decoder := json.NewDecoder(bytes.NewReader(contents))
	decoder.DisallowUnknownFields()
	var pages PageSet
	if err := decoder.Decode(&pages); err != nil {
		return PageSet{}, fmt.Errorf("awsmodel: parse coverage pages: %w", err)
	}
	if pages.Category == "" {
		return PageSet{}, fmt.Errorf("awsmodel: coverage pages: category is empty")
	}

	slugs := map[string]bool{}
	if err := validatePage("index", pages.Index, slugs); err != nil {
		return PageSet{}, err
	}
	// A known service may be left out. S3 is: Predastore serves that surface,
	// so the gateway dispatch tables hold no honest number to publish for it.
	for _, service := range Services() {
		page, ok := pages.Services[service]
		if !ok {
			continue
		}
		if err := validatePage(string(service), page, slugs); err != nil {
			return PageSet{}, err
		}
	}
	for service := range pages.Services {
		if _, ok := modelFiles[service]; !ok {
			return PageSet{}, fmt.Errorf("awsmodel: coverage pages: unknown service %q", service)
		}
	}
	return pages, nil
}

func validatePage(owner string, page PageMetadata, slugs map[string]bool) error {
	fail := func(format string, args ...any) error {
		return fmt.Errorf("awsmodel: coverage page %s: %s", owner, fmt.Sprintf(format, args...))
	}
	if !slugPattern.MatchString(page.Slug) {
		return fail("slug %q is not lowercase kebab-case", page.Slug)
	}
	if slugs[page.Slug] {
		return fail("slug %q is used by another page", page.Slug)
	}
	slugs[page.Slug] = true
	if page.Name == "" || page.Title == "" {
		return fail("name and title are required")
	}
	if len(page.Tags) == 0 {
		return fail("at least one tag is required")
	}
	if !strings.HasSuffix(page.SEOTitle, seoTitleSuffix) {
		return fail("seoTitle %q does not end with %q", page.SEOTitle, seoTitleSuffix)
	}
	if length := utf8.RuneCountInString(page.SEOTitle); length < seoTitleMin || length > seoTitleMax {
		return fail("seoTitle is %d characters, want %d-%d", length, seoTitleMin, seoTitleMax)
	}
	if length := utf8.RuneCountInString(page.Description); length < descriptionMin || length > descriptionMax {
		return fail("description is %d characters, want %d-%d", length, descriptionMin, descriptionMax)
	}
	return nil
}

// RenderServicePage renders one publishable coverage page. intro is optional
// hand-written prose emitted above the table; it may not open a new page
// section, so its headings must be level three or deeper.
func RenderServicePage(coverage OperationCoverage, pages PageSet, intro string) (string, error) {
	if err := checkIntro(string(coverage.Service), intro); err != nil {
		return "", err
	}
	page, ok := pages.Services[coverage.Service]
	if !ok {
		return "", fmt.Errorf("awsmodel: coverage pages: no entry for service %q", coverage.Service)
	}
	// An opaque service has no dispatch table to count, so there is no honest
	// number to put on a page. It is left out until a real one exists.
	if coverage.Opaque {
		return "", fmt.Errorf("awsmodel: %s coverage is not enumerable, so it has no publishable page", coverage.Service)
	}

	var body strings.Builder
	writeFrontmatter(&body, page, pages.Category)
	fmt.Fprintf(&body, "# %s\n\n## Overview\n\n", page.Title)

	fmt.Fprintf(&body, "Spinifex implements **%d of the %d** operations in the %s `%s` API model, as pinned in `aws-sdk-go %s` — **%.1f%%**.\n\n",
		len(coverage.Implemented), len(coverage.Modelled), page.Name, coverage.APIVersion, SourceSDKVersion, coverage.ImplementedPercent())

	if intro != "" {
		body.WriteString(strings.TrimSpace(intro) + "\n\n")
	}

	body.WriteString("### Operations\n\n| Operation | Status |\n|---|---|\n")
	for _, status := range coverage.OperationStatuses() {
		fmt.Fprintf(&body, "| `%s` | %s |\n", status.Operation, status.Status)
	}
	return body.String(), nil
}

// RenderIndexPage renders the cross-service summary that links to each page.
func RenderIndexPage(coverages []OperationCoverage, pages PageSet, intro string) (string, error) {
	if err := checkIntro("index", intro); err != nil {
		return "", err
	}

	var body strings.Builder
	writeFrontmatter(&body, pages.Index, pages.Category)
	fmt.Fprintf(&body, "# %s\n\n## Overview\n\n", pages.Index.Title)
	fmt.Fprintf(&body, "Spinifex serves the AWS APIs below. Every page counts the operations in the pinned `aws-sdk-go %s` `api-2.json` model for its service and reports, operation by operation, whether Spinifex implements it.\n\n", SourceSDKVersion)
	body.WriteString("| Service | API version | Implemented | Modelled | Coverage |\n|---|---|---:|---:|---:|\n")
	for _, coverage := range sortedCoverages(coverages) {
		page, ok := pages.Services[coverage.Service]
		if !ok {
			continue
		}
		fmt.Fprintf(&body, "| [%s](/docs/%s) | %s | %d | %d | %.1f%% |\n",
			page.Name, page.Slug, coverage.APIVersion, len(coverage.Implemented), len(coverage.Modelled), coverage.ImplementedPercent())
	}

	if intro != "" {
		body.WriteString("\n" + strings.TrimSpace(intro) + "\n")
	}
	return body.String(), nil
}

func writeFrontmatter(body *strings.Builder, page PageMetadata, category string) {
	body.WriteString("---\n")
	fmt.Fprintf(body, "title: %s\n", strconv.Quote(page.Title))
	fmt.Fprintf(body, "seoTitle: %s\n", strconv.Quote(page.SEOTitle))
	fmt.Fprintf(body, "description: %s\n", strconv.Quote(page.Description))
	fmt.Fprintf(body, "category: %s\n", strconv.Quote(category))
	body.WriteString("sections:\n  - overview\ntags:\n")
	for _, tag := range page.Tags {
		fmt.Fprintf(body, "  - %s\n", tag)
	}
	body.WriteString("---\n\n")
}

func checkIntro(owner, intro string) error {
	for line := range strings.SplitSeq(intro, "\n") {
		if strings.HasPrefix(line, "# ") || strings.HasPrefix(line, "## ") {
			return fmt.Errorf("awsmodel: %s intro heading %q opens a new page section; use level three or deeper", owner, strings.TrimSpace(line))
		}
	}
	return nil
}
