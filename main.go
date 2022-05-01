package main

import (
	"fmt"
	"go_practice/ch2/popcount"
	"go_practice/ch2/practice1"
	"go_practice/ch2/practice3"
	"go_practice/ch2/practice4"
	"go_practice/ch2/practice5"
	"go_practice/ch3/comma"
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
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline)
	complex := complex(1, -2)
	fmt.Println(complex)
	fmt.Println(real(complex))
	fmt.Println(imag(complex))
	fmt.Println(comma.Comma("1234567890"))
	fmt.Printf("%T, %[1]v\n", [...]int{1, 2, 3})
	fmt.Printf("%T, %[1]v\n", [...]int{10: -1})
}
