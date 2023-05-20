package main

import (
	"context"
	"fmt"
	"go_practice/mercari_expert/ch15/checkdigit"
	"strconv"
	"time"
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

	jan13 := checkdigit.NewJAN13()
	seed = "456995111617"
	digit, err = jan13.Generate(seed)
	if err != nil {
		fmt.Println(err)
	}
	code = seed + strconv.Itoa(digit)
	fmt.Println(seed, digit, code)
	fmt.Println(jan13.Verify(code))

	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, time.Second)
	go func() {
		for i := 0; ; i++ {
			fmt.Println(i)
		}
	}()
	for {
		select {
		case <-timeout.Done():
			return
		default:
			cancel()
		}
	}
}
