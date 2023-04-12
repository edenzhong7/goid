package goid

import (
	"github.com/fengyoulin/inspect"
	"reflect"
)

var offset uintptr

func init() {
	typ := inspect.TypeOf("runtime.g")
	if typ == nil {
		panic("runtime.g not found")
	}
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.Name == "goid" && (f.Type.Kind() == reflect.Int64 || f.Type.Kind() == reflect.Uint64) {
			offset = f.Offset
			break
		}
	}
	if offset == 0 {
		panic("runtime.g.goid not found")
	}
}
