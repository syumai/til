package main

import (
	"syscall/js"
	"uuid"
)

func generateUUIDv7(this js.Value, args []js.Value) any {
	return uuid.NewV7().String()
}

func main() {
	js.Global().Set("generateUUIDv7", js.FuncOf(generateUUIDv7))
	select {} // keep the Wasm instance alive
}
