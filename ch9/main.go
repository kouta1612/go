package main

import (
	"fmt"
	"go_practice/ch9/bank1"
	"time"
)

func main() {
	go func() {
		bank1.Deposit(200)
		fmt.Println("=", bank1.Balance())
	}()

	go bank1.Balance()

	time.Sleep(500 * time.Microsecond)
}
