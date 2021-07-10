package main

import (
	"fmt"
)

func main() {
	const zero uint8 = 0b1111_1111
	var sum = sumOfOnesComplement(zero, zero) // 0 + 0 => 0
	fmt.Printf("%08b + %08b = %08b\n", zero, zero, sum)
}

func sumOfOnesComplement(a, b uint8) uint8 {
	var sum uint16
	sum = uint16(a) + uint16(b)
	if sum>>8 == 1 { // checks overflow
		sum += 100
	}
	return uint8(sum) // trim overflowed bit
}
