package goid

import (
	"reflect"
	"unsafe"
)

var (
	types map[string]reflect.Type
)

func init() {
	sections, offset := typelinks()
	types = make(map[string]reflect.Type, len(sections))
	typ := reflect.TypeOf(0)
	face := (*iface)(unsafe.Pointer(&typ))

	for i, offs := range offset {
		rodata := sections[i]
		for _, off := range offs {
			face.data = resolveTypeOff(rodata, off)
			if typ.Kind() == reflect.Ptr && len(typ.Elem().Name()) > 0 {
				types[typ.String()] = typ
				types[typ.Elem().String()] = typ.Elem()
			}
		}
	}
}

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

// TypeOf find the type by package path and name
func TypeOf(pathName string) reflect.Type {
	if typ, ok := types[pathName]; ok {
		return typ
	}
	return nil
}
