package main

import (
	"fmt"
	"go_practice/ch5/links"
	"log"
	"os"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(string) []string, workList []string) {
	seen := make(map[string]bool)
	for len(workList) > 0 {
		items := workList
		workList = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				workList = append(workList, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
