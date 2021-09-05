package main

func Compare[T comparable] (a, b T) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func main() {
	println(Compare(1, 2))
	println(Compare("b", "a"))
}
