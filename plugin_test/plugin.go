package main

import (
	"fmt"
	"unsafe"

	"github.com/AminRezaei0x443/quickjs-go"
)

/*
#cgo CFLAGS: -I/mnt/d/Projects/IL-00/CO-443/quickjs-go/3rdparty/include/quickjs
#cgo linux,!android,amd64 LDFLAGS: -L/mnt/d/Projects/IL-00/CO-443/quickjs-go/3rdparty/libs/quickjs/Linux -lquickjs
#cgo linux,!android LDFLAGS: -lm -ldl -lpthread
#include "quickjs.h"

// extern JSModuleDef* js_init_module(JSContext* ctx, const char* module_name);

*/
import "C"

//export js_init_module
func js_init_module(ctx *C.JSContext, modName *C.char) *C.JSModuleDef {
	context := quickjs.WrapContext(unsafe.Pointer(ctx))
	defer context.FreeVals()
	moduleName := C.GoString(modName)
	m := context.DefineModule(moduleName, func(ctx *quickjs.Context, module *quickjs.Module) int {
		defer ctx.FreeVals()
		fmt.Println("Testing Module Def Call")
		module.AddProperty("dummyProp", ctx.Int32(2))
		module.AddFunction("dummyFunc", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
			return ctx.Int32(1)
		})
		return 1
	})
	m.ExportName("dummyProp")
	m.ExportName("dummyFunc")
	return (*C.JSModuleDef)(m.Ref())
}

// func main() {}
