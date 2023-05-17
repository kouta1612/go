package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()
	for {
		select {
		case <-time.Tick(1 * time.Second):
			fmt.Println("waiting...")
		case <-done:
			fmt.Println("done!")
			return
		}
	}
}
