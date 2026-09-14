package awsmodel_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	. "github.com/mulgadc/spinifex/internal/awsmodel"
)

// stagePages copies the checked-in metadata into a temp directory so a write
// test cannot touch docs/compatibility.
func stagePages(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	contents, err := os.ReadFile(filepath.Join("..", "..", "docs", "compatibility", "pages.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "pages.json"), contents, 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func stsCoverage(t *testing.T) OperationCoverage {
	t.Helper()
	coverage, err := CompareOperations(STS, DispatchInventory{Registered: []string{"AssumeRole"}})
	if err != nil {
		t.Fatal(err)
	}
	return coverage
}

func TestWritePagesWritesAServicePageAndTheIndex(t *testing.T) {
	dir := stagePages(t)

	if err := WritePages(dir, []OperationCoverage{stsCoverage(t)}); err != nil {
		t.Fatal(err)
	}

	service, err := os.ReadFile(filepath.Join(dir, "sts-api-coverage", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(service), "| `AssumeRole` | "+StatusImplemented+" |") {
		t.Errorf("service page is missing its operation table:\n%s", service)
	}
	index, err := os.ReadFile(filepath.Join(dir, "aws-api-coverage", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(index), "[STS](/docs/sts-api-coverage)") {
		t.Errorf("index does not link the service page:\n%s", index)
	}
}

func TestWritePagesIncludesAnIntroAndOverwritesTheGeneratedPage(t *testing.T) {
	dir := stagePages(t)
	pageDir := filepath.Join(dir, "sts-api-coverage")
	if err := os.MkdirAll(pageDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pageDir, "intro.md"), []byte("### Scope\n\nCarried prose.\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pageDir, "README.md"), []byte("stale hand edit"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := WritePages(dir, []OperationCoverage{stsCoverage(t)}); err != nil {
		t.Fatal(err)
	}

	page, err := os.ReadFile(filepath.Join(pageDir, "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(page), "stale hand edit") {
		t.Error("a hand edit to a generated page survived regeneration")
	}
	if !strings.Contains(string(page), "### Scope\n\nCarried prose.") {
		t.Errorf("page does not carry its intro:\n%s", page)
	}
}

// S3 has no page metadata, so it must be skipped rather than crashing the run
// or publishing a page with no honest number on it.
func TestWritePagesSkipsAServiceWithNoMetadata(t *testing.T) {
	dir := stagePages(t)
	coverage, err := CompareOperations(S3, DispatchInventory{Opaque: true, Note: "Served elsewhere."})
	if err != nil {
		t.Fatal(err)
	}

	if err := WritePages(dir, []OperationCoverage{coverage}); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(filepath.Join(dir, "s3-api-coverage")); !os.IsNotExist(err) {
		t.Errorf("a page was written for a service with no metadata: %v", err)
	}
}

func TestWritePagesReportsMissingMetadata(t *testing.T) {
	if err := WritePages(t.TempDir(), nil); err == nil {
		t.Fatal("WritePages accepted a directory with no pages.json")
	}
}
