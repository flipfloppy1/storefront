//go:build js && wasm

package js

import "syscall/js"

func SetInterval(function func(), intervalMS int) {
	js.Global().Call("setInterval", js.FuncOf(func(js.Value, []js.Value) any { function(); return nil }), intervalMS)
}

func SetTimeout(function func(), intervalMS int) {
	js.Global().Call("setTimeout", js.FuncOf(func(js.Value, []js.Value) any { function(); return nil }), intervalMS)
}
