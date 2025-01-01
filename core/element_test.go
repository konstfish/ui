package ui_test

import (
	"testing"

	ui "github.com/konstfish/ui/core"
)

func TestBasicTag(t *testing.T) {
	el := ui.NewElement("div").
		AddClass("foo").
		SetAttribute("id", "test")

	result, err := el.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `<div class="foo" id="test"></div>`
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestContentInTag(t *testing.T) {
	el := ui.NewElement("p").
		SetContent("Hello World")

	result, err := el.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<p>Hello World</p>"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestIdChange(t *testing.T) {
	el := ui.NewElement("div").
		SetId("test")

	el.SetId("test2")

	result, err := el.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `<div id="test2"></div>`
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestNestedComponents(t *testing.T) {
	el := ui.NewElement("div").
		AddChild(
			ui.NewElement("span").
				SetContent("Child"),
		)

	result, err := el.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<div><span>Child</span></div>"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestScriptInjection(t *testing.T) {
	el := ui.NewElement("div").
		SetContent("<script>alert('xss')</script>")

	result, err := el.Render()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "<div>&lt;script&gt;alert(&#39;xss&#39;)&lt;/script&gt;</div>"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}
