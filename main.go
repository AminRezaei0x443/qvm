package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/AminRezaei0x443/quickjs-go"
	"github.com/AminRezaei0x443/qvm/networking"
)

func check(err error) {
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(err)
	}
}

func main() {
	runtime := quickjs.NewRuntime()

	defer runtime.Free()
	context := runtime.NewContext()
	defer context.Free()

	networking.AddNetModule(context)

	if len(os.Args) > 1 {
		n := os.Args[1]
		context.EvaluateFile(n)
	}
	context.EvaluateFile("js_tests/test.js")
	fmt.Println()
}
