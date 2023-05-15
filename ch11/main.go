package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 擬似乱数生成器から24までのランダムな値を生成
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	fmt.Println(rng.Intn(25))
}
