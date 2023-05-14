package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println(struct{}{})
	fmt.Println(struct{ Name string }{Name: "kota"})

	aborted := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		aborted <- struct{}{}
	}()
loop:
	for {
		select {
		case <-aborted:
			fmt.Println("aborted!")
			break loop
		default:
			fmt.Println("doing.")
		}
	}

	var done = make(chan struct{})
	go func() {
		close(done)
	}()
	for {
		select {
		case <-done:
			fmt.Println("done!")
			return
		default:
			fmt.Println("default.")
		}
	}
}
