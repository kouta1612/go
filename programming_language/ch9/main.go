package main

import (
	"fmt"
	"programming_language/ch9/bank1"
	"programming_language/ch9/bank2"
	"programming_language/ch9/bank3"
	"sync"
	"time"
)

func main() {
	// bank関連の出力テスト
	testBank()

	// 出力のテスト
	print()
}

func print() {
	var wg sync.WaitGroup
	var x, y int

	wg.Add(1)
	go func() {
		defer wg.Done()
		x = 1
		fmt.Print("y:", y, " ")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		y = 1
		fmt.Print("x:", x, " ")
	}()
	wg.Wait()
	fmt.Println()
}

func testBank() {
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
