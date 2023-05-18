package main

import (
	"io"
	"math/rand"
	"os"
)

func main() {
	src := rand.NewSource(0)
	rand := rand.New(src)
	io.Copy(os.Stdout, rand)
}
