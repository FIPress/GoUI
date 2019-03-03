package goui

import "C"

// MenuType is an enum of menu type
type MenuType int

const (
	Container MenuType = iota //just a container item for sub items
	Custom
	Standard
	Separator
)

// MenuDef is to define a menu item
type MenuDef struct {
	Type     MenuType
	Title    string
	HotKey   string
	Action   string
	Handler  func()
	Children []MenuDef
}

var actionMap map[string]func()

//export menuClicked
func menuClicked(action string) {
	println("menu clicked", action)
	f := actionMap[action]
	if f != nil {
		f()
	}
}
