package main

import "github.com/fipress/GoUI"

func main() {
	//goui.Service("chat/:msg", chatService)
	menuDefs := []goui.MenuDef{
		{Title: "AppMenu", Type: goui.Container, Children: []goui.MenuDef{
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

	goui.CreateWithMenu(goui.Settings{
		Title:  "Text Editor",
		Top:    30,
		Left:   100,
		Width:  300,
		Height: 440,
	},
		menuDefs)
}
