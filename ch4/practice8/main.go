package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	typeCounts := make(map[string]int)
	invalid := 0 //不正な UTF-8文字の数
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, error を返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsNumber(r) {
			typeCounts["Number"]++
		}
		if unicode.IsLetter(r) {
			typeCounts["String"]++
		}
	}
	fmt.Print("type\tcount\n")
	for c, n := range typeCounts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
