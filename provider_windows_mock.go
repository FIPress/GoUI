// +build windows

package goui

/*
#include "provider.h"
*/
import "C"

// this file is just for debug or testing purpose
// when there is no need for a real windows provider, yet need one to get built

func cCreate(cs C.WindowSettings, cMenuDefs *C.MenuDef, count C.int) {
}

func cInvokeJS(js *C.char) {
}

func cExit() {
}
