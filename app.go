//go:build js && wasm

package main

import "syscall/js"

func getFramesJs(this js.Value, args []js.Value) any {
	in := make([]byte, 0)
	n := js.CopyBytesToGo(in, args[0])
	if n <= 0 {
		return js.Null()
	}

	frames := getFrames(in)
	framesOut := make(map[int]any)

	for i, f := range frames {
		framesOut[i] = f.asMap()
	}

	return js.ValueOf(framesOut)
}

func main() {
	done := make(chan struct{}, 0)

	js.Global().Set("wasmGetFrames", js.FuncOf(getFramesJs))

	<-done
}
