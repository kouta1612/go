package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("-------------img--------------")
	for _, link := range visit(nil, doc, "img", "src") {
		fmt.Println(link)
	}
	fmt.Println("-------------script--------------")
	for _, link := range visit(nil, doc, "script", "src") {
		fmt.Println(link)
	}
	fmt.Println("-------------style--------------")
	for _, link := range visit(nil, doc, "link", "href") {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node, node string, href string) []string {
	if n.Type == html.ElementNode && n.Data == node {
		for _, a := range n.Attr {
			if a.Key == href {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c, node, href)
	}
	return links
}
