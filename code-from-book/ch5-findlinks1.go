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

// 下面这个方法能且仅能抓取超链接类型标签<a>， 而且只获取链接属性
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

// 来自课本上面的习题，编写一个函数能够处理其它类型的资源文件
func visit_extend(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode {
		var key string
		switch node.Data {
		case "a":
			key = "href"
		case "img":
			key = "src"
		case "script":
			key = "src"
		case "link":
			key = "href"
		}
		for _, a := range node.Attr {
			if a.Key == key {
				links = append(links, a.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
