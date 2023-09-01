package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting ", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 向通道写入数据
	channel <- Page{URL: url, Size: len(body)}
}

// 通过通道获取函数值，通过结构体打包具有逻辑关联的一组数据
func main() {
	pages := make(chan Page)
	urls := []string{
		"https://example.com",
		"https://go.dev/learn/"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("URL: %s, Size: %d\n", page.URL, page.Size)
	}
}
