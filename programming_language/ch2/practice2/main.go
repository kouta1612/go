package main

import (
	"fmt"
	"os"
	"programming_language/ch2/tempconv"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		var t float64
		fmt.Print("温度を数値で入力してください：")
		fmt.Scan(&t)
		tempPrint(t)
	}
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		tempPrint(t)
	}
}

func tempPrint(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
