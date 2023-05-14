package main

import (
	"fmt"
	"go_practice/ch9/bank1"
	"go_practice/ch9/bank2"
	"go_practice/ch9/bank3"
	"time"
)

func main() {
	// bank1
	go func() {
		bank1.Deposit(200)
		fmt.Println("=", bank1.Balance())
	}()

	go bank1.Balance()

	time.Sleep(500 * time.Microsecond)

	// bank2
	go func() {
		bank2.Deposit(200)
		fmt.Println("=", bank2.Balance())
	}()

	go bank2.Balance()

	time.Sleep(500 * time.Microsecond)

	// bank3
	go func() {
		bank3.Deposit(200)
		fmt.Println("=", bank3.Balance())
	}()

	go bank3.Balance()

	time.Sleep(500 * time.Microsecond)
}
