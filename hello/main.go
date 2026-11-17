package main

import (
	"syscall"
	"unsafe"
)

func main() {
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBox := user32.NewProc("MessageBoxW")
	text, _ := syscall.UTF16PtrFromString("こんにちは")
	title, _ := syscall.UTF16PtrFromString("あいさつ")
	messageBox.Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), 0x40) // MB_ICONINFORMATION
}
