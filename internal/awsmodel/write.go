package awsmodel

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	pagesFile = "pages.json"
	introFile = "intro.md"
	pageFile  = "README.md"
)

// WritePages renders the publishable coverage pages into outputDir, which must
// already hold pages.json. A service without page metadata is skipped, and any
// README.md already there is overwritten: only pages.json and the intro.md
// files are hand-maintained.
func WritePages(outputDir string, coverages []OperationCoverage) error {
	contents, err := os.ReadFile(filepath.Join(outputDir, pagesFile))
	if err != nil {
		return fmt.Errorf("awsmodel: read coverage page metadata: %w", err)
	}
	pages, err := ParsePageSet(contents)
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
		body, err := RenderServicePage(coverage, pages, intro)
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
	body, err := RenderIndexPage(coverages, pages, intro)
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
		return "", fmt.Errorf("awsmodel: read %s intro: %w", slug, err)
	}
	return string(contents), nil
}

func writePage(outputDir, slug, body string) error {
	dir := filepath.Join(outputDir, slug)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("awsmodel: create %s page directory: %w", slug, err)
	}
	if err := os.WriteFile(filepath.Join(dir, pageFile), []byte(body), 0o600); err != nil {
		return fmt.Errorf("awsmodel: write %s page: %w", slug, err)
	}
	return nil
}
