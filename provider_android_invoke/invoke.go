// +build android
// +build arm 386 amd64 arm64

// Package invoke is for invoking main.main
//
package provider_android_invoke

// InvokeMain calls main.main by its address.
func InvokeMain(ptr uintptr)
