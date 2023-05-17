package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	len := len(s)
	for i := 0; i < len; i++ {
		if (i+2)%3 == 0 && len > 3 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234"))
}
