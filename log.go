// +build !android

package goui

import "C"
import "fmt"

//export goLog
func goLog(msg *C.char) {
	fmt.Println(C.GoString(msg))
}

func Log(args ...interface{}) {
	fmt.Print(args...)
}

func Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
