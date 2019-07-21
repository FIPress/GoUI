//+build windows

package goui

/*
#cgo LDFLAGS: -static provider_windows.dll

#include "provider_windows.h"
#include "provider.h"

extern void menuClicked(_GoString_ s);
extern void handleClientReq(const char* s);

void createApp(WindowSettings settings, MenuDef* menuDefs, int menuCount) {
	seLogger(&goLog);
	setHandleClientReq(&handleClientReq);
	//goUILog("settings.url:%s",settings.url);
	create(settings, menuDefs, menuCount);
}

void invokeScript(const char* js) {
	invokeJS(js);
}

void exitApp() {
	exitWebview();
}

*/
import "C"

func createApp(cs C.WindowSettings, cMenuDefs *C.MenuDef, count C.int) {
	C.createApp(cs, cMenuDefs, count)
}

func invokeAppJS(js *C.char) {
	C.invokeScript(js)
}

func exitApp() {
	C.exitApp()
}
