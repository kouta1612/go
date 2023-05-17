package main

import "fmt"

func main() {
	persons := map[string]int{"kota": 29}
	immutable(persons)
	fmt.Println(persons)
	mutable(persons)
	fmt.Println(persons)
}

// mapは参照型なのでアドレス値が渡される
// 初期化されて変数に代入されると新しいアドレス値が入るだけなので元のmapは更新されない
func immutable(m map[string]int) {
	m = map[string]int{"erina": 29}
}

// mapの値を直接更新しているので元のmapも更新される
func mutable(m map[string]int) {
	m["erina"] = 29
}
