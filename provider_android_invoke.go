// +build android
// +build arm 386 amd64 arm64

package goui

import "C"

//export invokeMain
func invokeMain(ptr uintptr) {
	Log("invoke main")
	invoke.InvokeMain(ptr)
}
