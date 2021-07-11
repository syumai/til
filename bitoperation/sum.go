package main

import (
	"fmt"
)

func main() {
	var a, b, sum uint8
	a = 0b1111_1110                 // -1
	b = 0b0000_0011                 // 3
	sum = sumOfOnesComplement(a, b) // 2
	fmt.Printf("%d + %d = %d\n", toI(a), toI(b), toI(sum))

	a = 0b1000_0000                 // -127
	b = 0b0001_1011                 // 27
	sum = sumOfOnesComplement(a, b) // -100
	fmt.Printf("%d + %d = %d\n", toI(a), toI(b), toI(sum))
}

func sumOfOnesComplement(a, b uint8) uint8 {
	var sum uint16
	sum = uint16(a) + uint16(b)
	return uint8(sum + sum>>8) // add overflowed bit and trim to 8bits
}

func toI(c uint8) int8 {
	if c>>7 == 0 {
		return int8(c)
	}
	return -1 * int8(^c)
}
