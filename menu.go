package goui

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
	Children []MenuDef
}

func AddMenu() {

}

func SetMenus(defs []MenuDef) {

}
