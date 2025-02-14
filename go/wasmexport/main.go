package main

//go:wasmexport add
func add(x, y int32) int32 {
	return x + y
}

func main() {
}
