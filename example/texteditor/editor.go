package main

import (
	"github.com/fipress/GoUI"
	"runtime"
)

func main() {
	var menuDefs []goui.MenuDef
	switch runtime.GOOS {
	case "darwin":
		menuDefs = []goui.MenuDef{{Title: "AppMenu", Type: goui.Container, Children: []goui.MenuDef{
			{Title: "About", Type: goui.Standard, Action: "about"},
			{Type: goui.Separator},
			{Title: "Hide", Type: goui.Standard, Action: "hide"},
			{Type: goui.Separator},
			{Title: "Quit", Type: goui.Standard, Action: "quit"},
		}},
			{Title: "File", Type: goui.Container, Children: []goui.MenuDef{
				{Title: "New", Type: goui.Custom, Action: "new", Handler: func() {
					println("new file")
				}},
				{Title: "Open", Type: goui.Custom, Action: "open"},
				{Type: goui.Separator},
				{Title: "Save", Type: goui.Custom, Action: "save"},
				{Title: "Save as", Type: goui.Custom, Action: "saveAs"},
			}},
			{Title: "Edit", Type: goui.Container, Children: []goui.MenuDef{
				{Title: "Undo", Type: goui.Custom, Action: "undo"},
				{Title: "Redo", Type: goui.Custom, Action: "redo"},
				{Type: goui.Separator},
				{Title: "Cut", Type: goui.Custom, Action: "cut"},
				{Title: "Copy", Type: goui.Custom, Action: "copy"},
				{Title: "Paste", Type: goui.Custom, Action: "paste"},
			}},
		}
	case "windows":
		println("windows")
	default:
		menuDefs = []goui.MenuDef{{Title: "File", Type: goui.Container, Children: []goui.MenuDef{
			{Title: "New", Type: goui.Custom, Action: "new", Handler: func() {
				println("new file")
			}},
			{Title: "Open", Type: goui.Custom, Action: "open"},
			{Type: goui.Separator},
			{Title: "Save", Type: goui.Custom, Action: "save"},
			{Title: "Save as", Type: goui.Custom, Action: "saveAs"},
			{Type: goui.Separator},
			{Title: "Quit", Type: goui.Standard, Action: "quit"},
		}},
			{Title: "Edit", Type: goui.Container, Children: []goui.MenuDef{
				{Title: "Undo", Type: goui.Custom, Action: "undo"},
				{Title: "Redo", Type: goui.Custom, Action: "redo"},
				{Type: goui.Separator},
				{Title: "Cut", Type: goui.Custom, Action: "cut"},
				{Title: "Copy", Type: goui.Custom, Action: "copy"},
				{Title: "Paste", Type: goui.Custom, Action: "paste"},
			}},
		}
	}

	goui.CreateWithMenu(goui.Settings{
		Title:  "Text Editor",
		Url:    "./ui/editor.html",
		Top:    30,
		Left:   100,
		Width:  430,
		Height: 450,
		Debug:  true,
	},
		menuDefs)
}
