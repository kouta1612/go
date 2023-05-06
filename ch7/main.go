package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Println(rw)

	w = nil
	rw, ok := w.(io.ReadWriter)
	fmt.Println(ok)

	_, err := os.Open("/no/such/file")
	fmt.Printf("%#v\n", err)
	fmt.Printf("%v\n", err)
	io.WriteString(os.Stdout, "Hello, World\n")

}
