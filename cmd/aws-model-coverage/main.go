package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mulgadc/spinifex/internal/awsmodel"
	"github.com/mulgadc/spinifex/spinifex/gateway"

	_ "github.com/mulgadc/bluebottle/pkg/fipsboot"
)

const (
	pagesFile = "pages.json"
	introFile = "intro.md"
	pageFile  = "README.md"
)

func main() {
	outputDir := flag.String("out", "", "write one publishable page per service into this directory")
	flag.Parse()

	coverages, err := compareAll()
	if err != nil {
		fail(err)
	}
	if *outputDir == "" {
		fmt.Print(awsmodel.RenderCoverageSummary(coverages))
		return
	}
	if err := writePages(*outputDir, coverages); err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "aws-model-coverage: %v\n", err)
	os.Exit(1)
}

func compareAll() ([]awsmodel.OperationCoverage, error) {
	dispatch := gateway.AWSOperationInventory()
	coverages := make([]awsmodel.OperationCoverage, 0, len(awsmodel.Services()))
	for _, service := range awsmodel.Services() {
		inventory, ok := dispatch[string(service)]
		var modelInventory awsmodel.DispatchInventory
		switch {
		case ok:
			modelInventory = awsmodel.DispatchInventory{
				Registered:  inventory.Registered,
				Stubbed:     inventory.Stubbed,
				Unsupported: inventory.Unsupported,
			}
		case service == awsmodel.S3:
			modelInventory = awsmodel.DispatchInventory{
				Opaque: true,
				Note:   "Spinifex delegates the S3 REST surface to Predastore, which has no operation-name dispatch table to compare mechanically.",
			}
		default:
			return nil, fmt.Errorf("no gateway inventory for %s", service)
		}

		coverage, err := awsmodel.CompareOperations(service, modelInventory)
		if err != nil {
			return nil, err
		}
		coverages = append(coverages, coverage)
	}
	return coverages, nil
}

func writePages(outputDir string, coverages []awsmodel.OperationCoverage) error {
	contents, err := os.ReadFile(filepath.Join(outputDir, pagesFile))
	if err != nil {
		return fmt.Errorf("read page metadata: %w", err)
	}
	pages, err := awsmodel.ParsePageSet(contents)
	if err != nil {
		return err
	}

	for _, coverage := range coverages {
		page, ok := pages.Services[coverage.Service]
		if !ok {
			continue
		}
		intro, err := readIntro(outputDir, page.Slug)
		if err != nil {
			return err
		}
		body, err := awsmodel.RenderServicePage(coverage, pages, intro)
		if err != nil {
			return err
		}
		if err := writePage(outputDir, page.Slug, body); err != nil {
			return err
		}
	}

	intro, err := readIntro(outputDir, pages.Index.Slug)
	if err != nil {
		return err
	}
	body, err := awsmodel.RenderIndexPage(coverages, pages, intro)
	if err != nil {
		return err
	}
	return writePage(outputDir, pages.Index.Slug, body)
}

// readIntro returns the hand-written prose for a page, or "" where the page has
// none. An intro is optional by design: most services have nothing to add.
func readIntro(outputDir, slug string) (string, error) {
	contents, err := os.ReadFile(filepath.Join(outputDir, slug, introFile))
	if errors.Is(err, fs.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("read %s intro: %w", slug, err)
	}
	return string(contents), nil
}

func writePage(outputDir, slug, body string) error {
	dir := filepath.Join(outputDir, slug)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("create %s page directory: %w", slug, err)
	}
	if err := os.WriteFile(filepath.Join(dir, pageFile), []byte(body), 0o644); err != nil {
		return fmt.Errorf("write %s page: %w", slug, err)
	}
	return nil
}
