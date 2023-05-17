package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v", err)
		os.Exit(1)
	}
	outline(doc)
}

func outline(n *html.Node) {
	if n.Type == html.ElementNode && n.Data != "script" && n.Data != "style" && n.Data != "noscript" {
		if child := n.FirstChild; child != nil && child.Type == html.TextNode && strings.TrimSpace(child.Data) != "" {
			fmt.Println("------------------------------------------")
			fmt.Printf("%s\n", strings.TrimSpace(child.Data))
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(c)
	}
}
