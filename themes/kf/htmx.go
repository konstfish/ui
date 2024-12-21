package kf

import (
	ui "github.com/konstfish/ui/core"
)

func Placeholder(path string) *ui.Element {
	return ui.NewElement("span").
		SetAttribute("hx-get", path).
		SetAttribute("hx-swap", "outerHTML").
		SetAttribute("hx-trigger", "load")
}
