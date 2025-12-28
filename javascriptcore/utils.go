package javascriptcore

import (
	"unsafe"
)

func cstring(s string) uintptr {
	if s == "" {
		return uintptr(0)
	}

	byteSlice := make([]byte, len(s)+1)
	copy(byteSlice, s)
	byteSlice[len(s)] = 0

	return uintptr(unsafe.Pointer(&byteSlice[0]))
}

func goString(ptr uintptr) string {
	if ptr == 0 {
		return ""
	}

	p := unsafe.Pointer(ptr)
	n := 0
	for {
		b := *(*byte)(unsafe.Add(p, n))
		if b == 0 {
			break
		}
		n++
	}

	return string(unsafe.Slice((*byte)(p), n))
}
