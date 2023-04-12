package goid

import "unsafe"

// ID returns current goroutine's runtime ID
func ID() int64 {
	return *(*int64)(unsafe.Pointer(getg() + offset))
}
