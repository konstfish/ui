package ui_test

import (
	"regexp"
	"testing"

	ui "github.com/konstfish/ui/core"
)

func TestNewPage(t *testing.T) {
	page := ui.NewPage()

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	patterns := []string{
		`<html\s+lang="en"`,
		`<meta\s+charset="UTF-8"`,
		`<meta\s+(?:content="width=device-width,\s*initial-scale=1.0"\s+name="viewport"|name="viewport"\s+content="width=device-width,\s*initial-scale=1.0")`,
	}

	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, result)
		if err != nil {
			t.Fatalf("regex error: %v", err)
		}
		if !matched {
			t.Errorf("expected pattern %q not found in result", pattern)
		}
	}
}

func TestPageSetTitle(t *testing.T) {
	page := ui.NewPage().SetTitle("Test Page")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<title>Test Page</title>`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("title tag not found in result")
	}
}

func TestPageAddMeta(t *testing.T) {
	page := ui.NewPage().AddMeta("keywords", "test,page,golang")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<meta\s+(?:content="test,page,golang"\s+name="keywords"|name="keywords"\s+content="test,page,golang")`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("meta keywords tag not found in result")
	}
}

func TestPageSetDescription(t *testing.T) {
	page := ui.NewPage().SetDescription("Test description")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<meta\s+(?:content="Test description"\s+name="description"|name="description"\s+content="Test description")`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("meta description tag not found in result")
	}
}

func TestPageAddLink(t *testing.T) {
	page := ui.NewPage().AddLink("stylesheet", "/style.css")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<link\s+(?:href="/style\.css"\s+rel="stylesheet"|rel="stylesheet"\s+href="/style\.css")`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("stylesheet link tag not found in result")
	}
}

func TestPageAddLinkWithType(t *testing.T) {
	page := ui.NewPage().AddLinkWithType("alternate", "/feed.xml", "application/rss+xml")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<link\s+(?:(?:href="/feed\.xml"|rel="alternate"|type="application/rss\+xml")\s*){3}`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("RSS link tag not found in result")
	}
}

func TestPageAddScript(t *testing.T) {
	page := ui.NewPage().AddScript("/script.js")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	pattern := `<script\s+src="/script\.js"></script>`
	matched, err := regexp.MatchString(pattern, result)
	if err != nil {
		t.Fatalf("regex error: %v", err)
	}
	if !matched {
		t.Errorf("script tag not found in result")
	}
}

func TestPageChaining(t *testing.T) {
	page := ui.NewPage().
		SetTitle("Test Page").
		SetDescription("Test description").
		AddScript("/script.js").
		AddStyleSheet("/style.css")

	result, err := page.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	patterns := []string{
		`<title>Test Page</title>`,
		`<meta\s+(?:content="Test description"\s+name="description"|name="description"\s+content="Test description")`,
		`<script\s+src="/script\.js"></script>`,
		`<link\s+(?:href="/style\.css"\s+rel="stylesheet"|rel="stylesheet"\s+href="/style\.css")`,
	}

	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, result)
		if err != nil {
			t.Fatalf("regex error: %v", err)
		}
		if !matched {
			t.Errorf("expected pattern %q not found in result", pattern)
		}
	}
}
