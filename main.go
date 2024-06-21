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
	// print(params[2])

	return js.ValueOf(method + "/ " + url)
}
