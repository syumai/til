package main

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

type P struct { x int; y int }

func Marshal[T any](dest io.Writer, v *T) error {
	ary := (*[unsafe.Sizeof(*v)]byte)(unsafe.Pointer(v))
	s := ary[:]
	rd := bytes.NewReader(s)
	_, err := io.Copy(dest, rd)
	return err
}

func Unmarshal[T any](src io.Reader, v *T) error {
	ary := (*[unsafe.Sizeof(*v)]byte)(unsafe.Pointer(v))
	s := ary[:]
	_, err := io.ReadFull(src, s)
	return err
}

func main() {
	p := P{
		x: 128,
		y: 256,
	}

	var buf bytes.Buffer
	_ = Marshal(&buf, &p)

	var decoded P
	_ = Unmarshal(&buf, &decoded)

	fmt.Println(decoded)
}

