package main

import "fmt"

func main() {
	natures := make(chan int)
	squares := make(chan int)

	go counter(natures)
	go squarer(squares, natures)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
