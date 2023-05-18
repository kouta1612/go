package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func main() {
	src := rand.NewSource(0)
	rand := rand.New(src)
	io.CopyN(os.Stdout, rand, 10)

	var buf bytes.Buffer
	buf.Write([]byte("hagiwara "))
	buf.Write([]byte("kota."))
	fmt.Println(buf.String())

	io.Copy(os.Stdout, os.Stdin)
}
