package goui

/*
#include "menu.h"
*/
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

func convertMenuDef(def MenuDef) (cMenuDef C.MenuDef) {
	cMenuDef = C.MenuDef{}
	cMenuDef.title = C.CString(def.Title)
	cMenuDef.action = C.CString(def.Action)
	cMenuDef.key = C.CString(def.HotKey)
	cMenuDef.menuType = C.MenuType(def.Type)
	cMenuDef.children, cMenuDef.childrenCount = convertMenuDefs(def.Children)

	return
}

func convertMenuDefs(defs []MenuDef) (array *C.MenuDef, count C.int) {
	l := len(defs)
	if l == 0 {
		return
	}

	count = C.int(l)

	array = C.allocMenuDefArray(count)
	for i := 0; i < l; i++ {
		cMenuDef := convertMenuDef(defs[i])
		C.addChildMenu(array, cMenuDef, C.int(i))
	}

	return
}

var actionMap map[string]func()

//export menuClicked
func menuClicked(action *C.char) {
	a := C.GoString(action)
	println("menu clicked", a)
	f := actionMap[a]
	if f != nil {
		f()
	}
}
