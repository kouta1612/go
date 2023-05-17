package main

import (
	"fmt"
	"io"
)

type neverEnding byte

func main() {
	r := io.LimitReader(neverEnding('x'), 10)
	b, _ := io.ReadAll(r)
	fmt.Println(string(b))
}

func (b neverEnding) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}
