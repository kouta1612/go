package main

import (
	"log"
	"time"
)

// 今回のdefer文の例ではtrace関数は遅延されず、中の関数値が遅延される
// deferの対象が中の関数値だから
func main() {
	defer trace("bigSlowOperation")()
	// ...大量の処理...
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s: ", msg)
	return func() { log.Printf("exit %s: (%s)", msg, time.Since(start)) }
}
