package main

import "fmt"

type (
	black	struct{}
	white	struct{}
	blue	struct{}
)

var (
	Black	black
	White	white
	Blue	blue
)

type Color interface {
	black | white | blue
}

func ColorCode[T Color](c T) string {
	switch (interface{})(c).(type) {
	case black:
		return "#000000"
	case white:
		return "#FFFFFF"
	case blue:
		return "#0000FF"
	default:
		panic("unexpected color")
	}
}

func main() {
	fmt.Println(ColorCode(Black))
	fmt.Println(ColorCode(White))
	fmt.Println(ColorCode(Blue))
}
