package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(stringJoin("hello", " ", "world", " ", "kota"))
}

func stringJoin(c string, strings ...string) string {
	buf := bytes.NewBufferString(c)
	for _, s := range strings {
		buf.WriteString(s)
	}
	return buf.String()
}
