package main

import (
	"fmt"
	"go_practice/mercari_expert/ch15/checkdigit"
	"strconv"
)

func main() {
	luhn := checkdigit.NewLuhn()
	seed := "54576238982341"
	digit, err := luhn.Generate(seed)
	if err != nil {
		fmt.Println(err)
	}
	code := seed + strconv.Itoa(digit)
	fmt.Println(seed, digit, code)
	fmt.Println(luhn.Verify(code))
}
