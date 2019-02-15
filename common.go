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
		C.int(boolToInt(settings.Minimizable)),
	}
}

func toMenuDefs(defs []MenuDef) {

}

func boolToInt(b bool) (i int) {
	if b {
		i = 1
	}
	return
}
