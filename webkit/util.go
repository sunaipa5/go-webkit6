package webkit

import (
	"strings"
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

func gostring(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	const maxLen = 4096
	str := unsafe.String((*byte)(p), maxLen)
	if idx := strings.IndexByte(str, 0); idx != -1 {
		return str[:idx]
	}
	return str
}
