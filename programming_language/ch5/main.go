package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(&slice[0])
	mutable(slice)
	fmt.Println(&slice[0])
}

func mutable(slice []int) {
	fmt.Println(&slice[0])
	slice = append(slice, 4)
	fmt.Println(&slice[0]) // 他のスライスと唯一異なるアドレスをもつスライス
}
