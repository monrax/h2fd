//go:build js && wasm

package main

import (
	"strconv"
	"syscall/js"
)

func getFramesJs(this js.Value, args []js.Value) any {
	in := make([]byte, 1024)
	n := js.CopyBytesToGo(in, args[0])
	if n <= 0 {
		println("error: no bytes were read")
		return js.Null()
	}

	frames := getFrames(in[:n])
	framesOut := make(map[string]any)

	for i, f := range frames {
		framesOut[strconv.Itoa(i)] = f.asMap()
	}

	return js.ValueOf(framesOut)
}

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("wasmGetFrames", js.FuncOf(getFramesJs))

	<-done
}
