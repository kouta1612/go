package main

import "fmt"

func main() {
	natures := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			natures <- i
		}
		close(natures)
	}()

	go func() {
		for x := range natures {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
