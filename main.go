package main

import (
	"fmt"
	"go_practice/ch2/popcount"
	"go_practice/ch2/practice1"
	"go_practice/ch2/practice3"
	"go_practice/ch2/practice4"
	"go_practice/ch2/practice5"
)

var num uint64 = 0b1111111111111111111111111111111111110011111111111111111111110011
var x byte = 0b1000

func main() {
	fmt.Println(practice1.CToK(0))
	fmt.Println(practice1.KToC(0))
	fmt.Println(practice1.KToF(0))
	fmt.Println(popcount.PopCount(num))
	fmt.Printf("%b\n", byte(num>>(8*7)))
	fmt.Println(practice3.PopCount(num))
	fmt.Println(practice4.PopCount(num))
	fmt.Println(x)
	fmt.Println(x & (x - 1))
	fmt.Println(practice5.PopCount(num))
}
