package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// (base) .../MyCode/go$ go build  -o fetch   ./code-from-book/ch1-fetch.go
// (base) .../MyCode/go$ go build  -o outline ./code-from-book/ch5-outline.go
// (base) .../MyCode/go$ ./fetch http://gopl.io | ./outline
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	// 获取页面大纲
	outline(nil, doc)

	// 统计各种节点的类型
	hash := make(map[string]int)
	countNode(hash, doc)
	for elem, n := range hash {
		fmt.Printf("Number of %6sv is  %3d\n", elem, n)
	}

	texts := getText(nil, doc)
	for _, txt := range texts {
		fmt.Println(txt)
	}
}

func outline(stack []string, node *html.Node) {
	if node.Type == html.ElementNode {
		stack = append(stack, node.Data)
		fmt.Println(stack)
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

/*
	 来自教材上面的练习, 编写一个函数统计HTML文本各种类型的节点个数
			若是 html.ElementNode	node.Data 便是标签元素类型
			若是 html.TextNode		node.Data 便是标签内部文本
			...
*/
func countNode(hash map[string]int, node *html.Node) {
	if node.Type == html.ElementNode {
		hash[node.Data]++
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		countNode(hash, c)
	}
}

// 来自教材上面的练习，编写一个函数获取所有文本节点的文本，但是忽略script和style两个标签
func getText(texts []string, node *html.Node) []string {
	if node.Type == html.TextNode {
		parent := node.Parent
		if parent != nil && (parent.Data == "script" || parent.Data == "style") {
			return texts
		}
		texts = append(texts, strings.Trim(node.Data, " \n\t\r\f\v"))
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		texts = getText(texts, c)
	}
	return texts
}
