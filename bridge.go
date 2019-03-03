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
	/*
		count := len(def.Children)
		ma := C.newMenuArray(C.int(count))
		for i:=0;i<count;i++ {
			cMenuDef := convertMenuDef(def.Children[i])
			C.addChildMenu(ma,cMenuDef,C.int(i))
		}*/

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

/*
func allocateArray(count int) []C.MenuDef {
	return C.malloc(C.sizeof_MenuDef*count);
}

func addMenu(array *C.MenuDef,def MenuDef,index int)  {
	cDef,childrenCount := convertMenuDef(def)
	//C.addChild(array,cDef,index)
	array[index] = cDef;
	for i:=0; i<childrenCount;i++ {
		addMenu(cDef.children,def.Children[i],i)
	}
}

func convertMenuDef(def MenuDef) (menuDef C.MenuDef, childrenCount int) {
	cTitle := C.CString(def.Title)
	defer C.free(unsafe.Pointer(cTitle))

	cAction := C.CString(def.Action)
	defer C.free(unsafe.Pointer(cAction))

	cKey :=	C.CString(def.HotKey)
	defer C.free(unsafe.Pointer(cKey))

	childrenCount = len(def.Children)

	menuDef = C.newMenuDef(C.MenuType(def.Type),cTitle,cAction,cKey,C.int(childrenCount))
	return
}*/
/*
func convertMenuDefs(defs []MenuDef) (menuDefs *C.MenuDef, size int) {
	size = len(defs)
	if size == 0 {
		return
	}

	//type_a *p = (type_a*)malloc(sizeof(type_a)+100*sizeof(int));
	menuDefs = C.malloc(C.sizeof_MenuDef*size)
	for i := 0; i < size; i++ {
		menuDefs[i] = convertMenuDef(defs[i])
	}
	return
}*/

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}
