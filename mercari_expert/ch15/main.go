package main

import (
	"fmt"
	"go_practice/mercari_expert/ch15/checkdigit"
	"strconv"
)

func main() {
	luhn := checkdigit.NewLuhn()
	seed := "545762389823411"
	digit, _ := luhn.Generate(seed)
	code := seed + strconv.Itoa(digit)
	fmt.Println(seed, digit, code)
	fmt.Println(luhn.Verify(code))
}
