package kf

import ui "github.com/konstfish/ui/core"

func Spinner(text string) *ui.Element {
	spinner := ui.NewElement("span")

	spinner.AddChild(ui.NewElement("span").AddClass("spinner"))

	if text != "" {
		spinner.SetContent(" " + text)
	}

	return spinner
}
