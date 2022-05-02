package main

import "fmt"

func main() {
	datas := [10]int{1, 2, 3, 4}
	reverse(&datas)
	fmt.Printf("%v\n", datas)
}

func reverse(a *[10]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
