package networking

import (
	"errors"
	"io/ioutil"
	"net/http"
	"unsafe"

	"github.com/AminRezaei0x443/quickjs-go"
	"github.com/AminRezaei0x443/qvm/core"
)

func InitHttpClass(ctx *quickjs.Context) quickjs.Value {
	var httpCls *quickjs.Class
	httpCls = ctx.NewClass("HttpClient", func(rt *quickjs.Runtime, val quickjs.Value) {
		so := (*core.StoredObject)(val.GetOpaque(httpCls.Id()))
		so.Remove()
	})
	clsObj := httpCls.DefConstructor(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		if len(args) < 2 {
			return ctx.ThrowError(errors.New("args < 2"))
		}
		method := args[0].String()
		url := args[1].String()
		req, e := http.NewRequest(method, url, nil)
		if e != nil {
			return ctx.ThrowError(e)
		}
		obj := httpCls.NewObject()
		stored := core.AppendObject(req)
		obj.SetOpaque(unsafe.Pointer(stored))
		return obj
	})
	httpCls.Proto().SetFunction("download", func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		so := (*core.StoredObject)(this.GetOpaque(httpCls.Id()))
		req := so.Get().(*http.Request)
		client := http.Client{}
		resp, e := client.Do(req)
		if e != nil {
			return ctx.ThrowError(e)
		}
		bytes, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			return ctx.ThrowError(e)
		}
		return ctx.String(string(bytes))
	})
	httpCls.ProtoStabilize()
	return clsObj
}
