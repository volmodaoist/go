package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

/*
type NodeType int32
type Attribute struct{
	Key, Val string
}

type Node struct{
	Type 	NodeType
	Data	string
	Attr	[]Attribute
	FirstChild, NextSibling *Node
}
*/

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "finklinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
