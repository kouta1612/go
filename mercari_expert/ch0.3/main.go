package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int, done <-chan struct{}) {
			for {
				select {
				case <-done:
					log.Printf("finished %d\n", i)
					return
				default:
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i, done)
	}

	close(done)
	time.Sleep(300 * time.Millisecond)
}
