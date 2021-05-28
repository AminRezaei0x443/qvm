package networking

import (
	"github.com/AminRezaei0x443/quickjs-go"
)

func AddNetModule(context *quickjs.Context) *quickjs.Module {
	m := context.DefineModule("networking", func(ctx *quickjs.Context, module *quickjs.Module) int {
		net := InitHttpClass(ctx)
		module.AddProperty("HttpClient", net)
		return 1
	})
	m.ExportName("HttpClient")
	return m
}
