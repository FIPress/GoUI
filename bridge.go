package goui

/*
#include <stdlib.h>
#include "bridge.c"

*/
import "C"
import (
	"path"
)

func convertSettings(settings Settings) C.WindowSettings {
	dir := path.Dir(settings.Url)

	return C.WindowSettings{C.CString(settings.Title),
		C.CString(settings.Url),
		C.CString(dir),
		C.int(settings.Left),
		C.int(settings.Top),
		C.int(settings.Width),
		C.int(settings.Height),
		C.int(boolToInt(settings.Resizable)),
		C.int(boolToInt(settings.Debug)),
	}
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
	len := len(defs)
	if len == 0 {
		return
	}

	count = C.int(len)

	array = C.allocMenuDefArray(count)
	for i := 0; i < len; i++ {
		cMenuDef := convertMenuDef(defs[i])
		C.addChildMenu(array, cMenuDef, C.int(i))
	}

	return
}

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}
