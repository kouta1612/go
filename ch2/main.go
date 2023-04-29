package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "国"
	// バイト数を返す
	fmt.Println(len(s))
	// ルーン数を返す
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println([]byte(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
