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
		for {
			x, ok := <-natures
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
