package main

import (
	"crypto/sha256"
	"fmt"
	"math"
)

func popCount(bytes [32]byte) int {
	result := 0
	for i := 0; i < 32; i++ {
		for bytes[i] != 0 {
			bytes[i] = bytes[i] & (bytes[i] - 1)
			result++
		}
	}
	return result
}

func diffPopCount(bytes1 [32]byte, bytes2 [32]byte) int {
	count1 := popCount(bytes1)
	count2 := popCount(bytes2)
	return int(math.Abs(float64(count1 - count2)))
}

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("c1: %b\nc2: %b\n", c1, c2)
	fmt.Println(popCount(c1))
	fmt.Println(popCount(c2))
	fmt.Println(popCount([32]byte{0b10000001, 0b00011001}))
	fmt.Println(diffPopCount(c1, c2))
}
