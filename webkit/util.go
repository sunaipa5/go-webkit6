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
