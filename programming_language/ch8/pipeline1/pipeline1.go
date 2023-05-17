package main

import "fmt"

func main() {
	natures := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; ; i++ {
			natures <- i
		}
	}()

	go func() {
		for {
			x := <-natures
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}
