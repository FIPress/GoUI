package goui

/*
#include <stdlib.h>
#include "provider.h"

*/
import "C"
import (
	"os"
	"path"
	"runtime"
	"unsafe"
)

const defaultDir = "web"
const defaultIndex = "index.html"

func convertSettings(settings Settings) C.WindowSettings {
	//dir := path.Dir(settings.Url)
	if settings.WebDir == "" {
		settings.WebDir = defaultDir
	}

	if settings.Index == "" {
		settings.Index = defaultIndex
	}

	if settings.Url == "" {
		settings.Url = path.Join(settings.WebDir, settings.Index)
		if runtime.GOOS == "linux" {
			wd, _ := os.Getwd()
			settings.Url = path.Join("file://", wd, settings.Url)
		} else if runtime.GOOS == "android" {
			settings.Url = path.Join("file:///android_asset/", settings.Url)
		}
	}

	// windows needs WebDir and Index
	// macOS and iOS need Url

	return C.WindowSettings{C.CString(settings.Title),
		C.CString(settings.WebDir),
		//C.CString(abs),
		C.CString(settings.Index),
		C.CString(settings.Url),
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

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}

func create(settings Settings, menuDefs []MenuDef) {
	//C.Create((*C.WindowSettings)(unsafe.Pointer(settings)))
	cs := convertSettings(settings)
	cMenuDefs, count := convertMenuDefs(menuDefs)
	cCreate(cs, cMenuDefs, count)
}

func activate() {

}

func invokeJS(js string) {
	cJs := C.CString(js)
	Log("invoke:", js)
	defer C.free(unsafe.Pointer(cJs))
	cInvokeJS(cJs)
}

func exit() {
	cExit()
}
