package main

import "fmt"

type T struct {
	Number int
}

func main() {
	s := []T{{1}, {2}, {3}}
	for _, v := range s {
		// 変数vは使い回されるので同じアドレスを参照
		fmt.Printf("%p\n", &v)
	}

	p := []*T{{1}, {2}, {3}}
	for _, v := range p {
		fmt.Printf("%p\n", v)
	}
}
