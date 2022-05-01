package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	var str string
	fmt.Print("ハッシュに変換したい文字列を入力してください: ")
	fmt.Scan(&str)
	f := flag.String("h", "sha256", "hash string")
	flag.Parse()
	if *f == "sha256" {
		fmt.Println(sha256.Sum256([]byte(str)))
	} else if *f == "sha384" {
		fmt.Println(sha512.Sum384([]byte(str)))
	} else if *f == "sha512" {
		fmt.Println(sha512.Sum512([]byte(*f)))
	} else {
		fmt.Println("invalid argument")
	}
}
