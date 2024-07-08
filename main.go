package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("request", js.FuncOf(request))
	<-make(chan bool)
}

func request(this js.Value, params []js.Value) interface{} {
	url := params[0].String()
	method := params[1].String()

	fetch := js.Global().Get("fetch")

	fetch.Call("apply", js.Undefined(), js.ValueOf([]interface{}{url, method})).Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		response := args[0]
		return response.Call("json").Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			data := args[0]
			js.Global().Get("console").Call("log", data)
			return nil
		}))
	})).Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		err := args[0]
		js.Global().Get("console").Call("log", err)
		return nil
	}))

	return nil
}
