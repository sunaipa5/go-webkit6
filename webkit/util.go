package webkit

import (
	"syscall"
	"unsafe"
)

type GList struct {
	Data uintptr
	Next *GList
	Prev *GList
}

func (x *GList) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

func cstring(s string) uintptr {
	ptr, err := syscall.BytePtrFromString(s)
	if err != nil {
		panic(err)
	}
	return uintptr(unsafe.Pointer(ptr))
}
