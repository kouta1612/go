package main

import (
	"fmt"
	"time"

	"programming_language/ch2/popcount"
	"programming_language/ch2/practice1"
	"programming_language/ch2/practice3"
	"programming_language/ch2/practice4"
	"programming_language/ch2/practice5"
	"programming_language/ch3/comma"
	"programming_language/ch6/geometry"
)

var num uint64 = 0b1111111111111111111111111111111111110011111111111111111111110011
var x byte = 3

func main() {
	fmt.Println(string('x' - 'a' + 'A'))
	fmt.Println(byte(num >> (x * 8)))
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
	fmt.Printf("%q\n", []rune("Hello, 世界"))
	fmt.Printf("%T\n%1T\n%1v\n", []int{}, [3]int{}, len(map[string]int{"alice": 32, "bob": 18}))
	msg := fmt.Sprintf("%q", []string{"aaa", "bbb"})
	maped := map[string]int{msg: 1}
	fmt.Println(maped)
	fmt.Printf("%v\n", []rune("a"))
	const day = 24 * time.Hour
	fmt.Println(day.Seconds())

	p, q := geometry.Point{X: 1, Y: 2}, geometry.Point{X: 4, Y: 6}
	fmt.Println(geometry.Distance(p, q))
}
