package goui

/*
#include <stdlib.h>
#include "common.c"
*/
import "C"
import (
	"path"
)

func toCSettings(settings Settings) C.WindowSettings {
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

func toMenuDef(def MenuDef) (menuDef C.MenuDef) {
	children, size := toMenuDefs(def.Children)
	return C.MenuDef{
		C.int(def.Type),
		C.CString(def.Title),
		C.CString(def.Action),
		C.CString(def.HotKey),
		children,
		C.int(size),
	}
}

func toMenuDefs(defs []MenuDef) (menuDefs *C.MenuDef, size int) {
	size = len(defs)
	if size == 0 {
		return
	}

	menuDefs = [size]C.MenuDef{}
	for i := 0; i < size; i++ {
		menuDefs[i] = toMenuDef(defs[i])
	}
	return
}

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}
