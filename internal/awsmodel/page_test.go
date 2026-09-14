package awsmodel_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/mulgadc/spinifex/internal/awsmodel"
)

func testPageSet(t *testing.T) PageSet {
	t.Helper()
	contents, err := os.ReadFile(filepath.Join("..", "..", "docs", "compatibility", "pages.json"))
	if err != nil {
		t.Fatal(err)
	}
	pages, err := ParsePageSet(contents)
	if err != nil {
		t.Fatal(err)
	}
	return pages
}

func TestParsePageSetAcceptsTheCheckedInFile(t *testing.T) {
	pages := testPageSet(t)

	if len(pages.Services) == 0 {
		t.Fatal("no service pages are configured")
	}
	// S3 is deliberately unpublished: Predastore serves that surface, so the
	// gateway dispatch tables hold no honest number for it.
	if _, ok := pages.Services[S3]; ok {
		t.Error("S3 has page metadata but no enumerable coverage to publish")
	}
}

func TestParsePageSetRejectsUnpublishableMetadata(t *testing.T) {
	for name, contents := range map[string]string{
		"unknown service":  `{"category":"Coverage","index":{},"services":{"notaservice":{}}}`,
		"missing category": `{"index":{},"services":{}}`,
		"unknown field":    `{"category":"Coverage","colour":"blue","index":{},"services":{}}`,
		"short seoTitle": `{"category":"Coverage","index":{"slug":"aws-api-coverage","name":"All","title":"All",` +
			`"seoTitle":"Too short — Spinifex Docs","description":"` + strings.Repeat("x", 155) + `","tags":["aws"]}}`,
		"missing suffix": `{"category":"Coverage","index":{"slug":"aws-api-coverage","name":"All","title":"All",` +
			`"seoTitle":"` + strings.Repeat("x", 55) + `","description":"` + strings.Repeat("x", 155) + `","tags":["aws"]}}`,
		"short description": `{"category":"Coverage","index":{"slug":"aws-api-coverage","name":"All","title":"All",` +
			`"seoTitle":"` + strings.Repeat("x", 39) + ` — Spinifex Docs","description":"too short","tags":["aws"]}}`,
		"bad slug": `{"category":"Coverage","index":{"slug":"AWS Coverage","name":"All","title":"All",` +
			`"seoTitle":"` + strings.Repeat("x", 39) + ` — Spinifex Docs","description":"` + strings.Repeat("x", 155) + `","tags":["aws"]}}`,
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := ParsePageSet([]byte(contents)); err == nil {
				t.Fatal("ParsePageSet accepted unpublishable metadata")
			}
		})
	}
}

func TestRenderServicePageCarriesEveryStatus(t *testing.T) {
	pages := testPageSet(t)
	coverage, err := CompareOperations(STS, DispatchInventory{
		Registered:  []string{"AssumeRole", "GetCallerIdentity", "GetSessionToken", "PublishInternal"},
		Stubbed:     []string{"GetSessionToken"},
		Unsupported: []string{"GetCallerIdentity"},
	})
	if err != nil {
		t.Fatal(err)
	}

	page, err := RenderServicePage(coverage, pages, "### Scope\n\nHand-written prose.\n")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		`title: "STS API Coverage"`,
		`category: "Coverage"`,
		"sections:\n  - overview\n",
		"# STS API Coverage\n\n## Overview\n",
		"Spinifex implements **1 of the 8** operations (12.5%) in the STS `2011-06-15` API model.",
		"### Scope\n\nHand-written prose.",
		"| `AssumeRole` | " + StatusImplemented + " |",
		"| `GetSessionToken` | " + StatusStub + " |",
		"| `GetCallerIdentity` | " + StatusNotApplicable + " |",
		"| `AssumeRoleWithSAML` | " + StatusNotImplemented + " |",
		"| `PublishInternal` | " + StatusOutsideModel + " |",
	} {
		if !strings.Contains(page, want) {
			t.Errorf("page does not contain %q:\n%s", want, page)
		}
	}
	if strings.Contains(page, "<details") || strings.Contains(page, "<summary") {
		t.Errorf("page contains raw HTML the docs site would escape:\n%s", page)
	}
}

// A modelled operation is never summarised as a count: a visitor asking whether
// Spinifex runs their workload is asking about the absent operations too.
func TestRenderServicePageRowsEveryModelledOperation(t *testing.T) {
	pages := testPageSet(t)
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}

	page, err := RenderServicePage(coverage, pages, "")
	if err != nil {
		t.Fatal(err)
	}
	for _, operation := range coverage.Modelled {
		if !strings.Contains(page, "| `"+operation+"` |") {
			t.Errorf("page has no row for modelled operation %q", operation)
		}
	}
}

func TestRenderServicePageWithoutIntro(t *testing.T) {
	pages := testPageSet(t)
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}

	page, err := RenderServicePage(coverage, pages, "")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(page, "\n\n\n") {
		t.Errorf("page has a blank run where the intro would be:\n%s", page)
	}
}

func TestRenderServicePageRejectsIntroOpeningASection(t *testing.T) {
	pages := testPageSet(t)
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := RenderServicePage(coverage, pages, "## Troubleshooting\n"); err == nil {
		t.Fatal("RenderServicePage accepted an intro that opens a new page section")
	}
}

// A "never" claim is only worth publishing if it cannot quietly become false.
func TestOperationStatusesRejectsStaleNotApplicableClaims(t *testing.T) {
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}

	for name, declared := range map[string]map[string]string{
		"not in the model": {"NotAnOperation": "a reason"},
		"no reason":        {"GetFederationToken": ""},
		"already implemented": {
			"AssumeRole": "claims the platform will never serve an operation it already serves",
		},
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := coverage.OperationStatuses(declared); err == nil {
				t.Fatal("OperationStatuses accepted a claim that cannot hold")
			}
		})
	}
}

func TestOperationStatusesPublishesDeclaredReasons(t *testing.T) {
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}

	statuses, err := coverage.OperationStatuses(map[string]string{"GetFederationToken": "No federation broker."})
	if err != nil {
		t.Fatal(err)
	}
	for _, status := range statuses {
		if status.Operation != "GetFederationToken" {
			continue
		}
		if status.Status != StatusNotApplicable || status.Note != "No federation broker." {
			t.Fatalf("status = %q, note = %q", status.Status, status.Note)
		}
		return
	}
	t.Fatal("declared operation is missing from the table")
}

// An opaque service has no dispatch table to count, so publishing a page for it
// would report 0% where the truth is that the surface is served elsewhere.
func TestRenderServicePageRefusesOpaqueCoverage(t *testing.T) {
	pages := testPageSet(t)
	coverage, err := CompareOperations(STS, DispatchInventory{
		Opaque: true,
		Note:   "Served elsewhere.",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := RenderServicePage(coverage, pages, ""); err == nil {
		t.Fatal("RenderServicePage published a page for a service with no enumerable coverage")
	}
}

func TestRenderIndexPageLinksEveryPublishedService(t *testing.T) {
	pages := testPageSet(t)
	coverages := make([]OperationCoverage, 0, len(Services()))
	for _, service := range Services() {
		coverage, err := CompareOperations(service, DispatchInventory{})
		if err != nil {
			t.Fatal(err)
		}
		coverages = append(coverages, coverage)
	}

	index, err := RenderIndexPage(coverages, pages, "")
	if err != nil {
		t.Fatal(err)
	}
	for _, service := range Services() {
		page, published := pages.Services[service]
		if !published {
			if strings.Contains(index, "/docs/"+string(service)+"-api-coverage") {
				t.Errorf("index links unpublished service %q", service)
			}
			continue
		}
		if !strings.Contains(index, "["+page.Name+"](/docs/"+page.Slug+")") {
			t.Errorf("index does not link %q at /docs/%s", page.Name, page.Slug)
		}
	}
}
