package webkit

import (
	"strings"
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
	if s == "" {
		return uintptr(0)
	}

	byteSlice := make([]byte, len(s)+1)
	copy(byteSlice, s)
	byteSlice[len(s)] = 0

	return uintptr(unsafe.Pointer(&byteSlice[0]))
}

func gostring(p uintptr) string {
	if p == 0 {
		return ""
	}
	s := unsafe.String((*byte)(unsafe.Pointer(p)), 4096)

	if idx := strings.IndexByte(s, 0); idx != -1 {
		return s[:idx]
	}
	return s
}
