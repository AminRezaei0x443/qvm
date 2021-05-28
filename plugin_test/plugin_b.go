package main

import (
	"fmt"
	"unsafe"

	"github.com/AminRezaei0x443/quickjs-go"
)

/*
#cgo linux,!android LDFLAGS: -lm -ldl -lpthread
*/
import "C"

//export js_init_module
func js_init_module(ctx unsafe.Pointer, modName *C.char) unsafe.Pointer {
	context := quickjs.WrapContext(ctx)
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
	return m.Ref()
}

func main() {}
