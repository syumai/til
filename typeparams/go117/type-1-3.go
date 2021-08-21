package main

type anySlice[T any] []T // constraintとしてanyが使えるか確認
type comparableSlice[T comparable] []T // constraintとしてanyが使えるか確認

func main() {}
