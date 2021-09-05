package main

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func NewMyOrdered[T Ordered](v T) MyOrdered[T] {
	return MyOrdered[T](v)
}

type MyOrdered[T Ordered] T

func (l MyOrdered[T]) Less(other T) bool {
	return T(l) < other
}

func main() {
	i := 1
	println(NewMyOrdered(i).Less(2))
}
