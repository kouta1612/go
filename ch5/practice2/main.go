package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
		os.Exit(1)
	}
	for node, count := range outline(make(map[string]int), doc) {
		fmt.Printf("%s: %d\n", node, count)
	}
}

func outline(mapping map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		mapping[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(mapping, c)
	}
	return mapping
}
