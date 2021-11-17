package main

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

type P struct {
	x int
	y int
}

type PEncoder struct {
	w io.Writer
}

func (enc *PEncoder) Encode(p P) error {
	size := unsafe.Sizeof(p)
	ptr := uintptr(unsafe.Pointer(&p))
	for i := uintptr(0); i < size; i++ {
		b := *(*byte)(unsafe.Pointer(ptr + i))
		if _, err := enc.w.Write([]byte{b}); err != nil {
			return err
		}
	}
	return nil
}

type PDecoder struct {
	r io.Reader
}

func (dec *PDecoder) Decode(p *P) error {
	size := unsafe.Sizeof(*p)
	ptr := uintptr(unsafe.Pointer(p))
	buf := make([]byte, 1)
	for i := uintptr(0); i < size; i++ {
		if _, err := dec.r.Read(buf); err != nil {
			return err
		}
		*(*byte)(unsafe.Pointer(ptr + i)) = buf[0]
	}
	return nil
}

func main() {
	p := P{
		x: 128,
		y: 256,
	}
	var buf bytes.Buffer
	enc := PEncoder{w: &buf}
	if err := enc.Encode(p); err != nil {
		panic(err)
	}

	var decoded P
	dec := PDecoder{r: &buf}
	if err := dec.Decode(&decoded); err != nil {
		panic(err)
	}

	fmt.Println(decoded)
}
