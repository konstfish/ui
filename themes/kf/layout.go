package kf

import ui "github.com/konstfish/ui/core"

func HorizontalRule() *ui.Element {
	return ui.NewElement("hr")
}

func VerticalRule() *ui.Element {
	return ui.NewElement("vr")
}

func Header(logo *ui.Element, links []KeyValue) *ui.Element {
	var headerContent *ui.Element = ui.NewElement("div").AddClass("content").AddChild(logo)

	var linksElement *ui.Element
	if links != nil {
		linksElement = ui.NewElement("nav")
		for _, kv := range links {
			linksElement.AddChild(Link(kv.Key, kv.Value))
		}

		headerContent.AddChild(linksElement)
	}

	return ui.NewElement("header").AddChild(headerContent)
}

func AppBody() *ui.Element {
	return ui.NewElement("div").AddId("app")
}

// FooterSimple creates a footer with text content.
func FooterSimple(content string) *ui.Element {
	return ui.NewElement("footer").AddChild(Text(content))
}
