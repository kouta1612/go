package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

type neverEnding byte

func main() {
	r := io.LimitReader(neverEnding('x'), 10)
	b, _ := io.ReadAll(r)
	fmt.Println(string(b))

	str := "kota"
	var buf bytes.Buffer
	io.Copy(&buf, strings.NewReader(str))
	fmt.Println(buf.String())

	f, _ := os.Create("sample.txt")
	defer f.Close()
	h := sha256.New()
	w := io.MultiWriter(f, h)
	io.WriteString(w, "hagiwara")
	fmt.Printf("%x\n", h.Sum(nil))
}

func (b neverEnding) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = byte(b)
	}
	return len(p), nil
}
