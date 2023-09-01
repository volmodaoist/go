package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// main 函数也是一个协程
func main() {
	go responseSize("https://example.com/")
	go responseSize("https://go.dev/learn/")
	// 如果不添加下面的等待语句，main函数可能会先于上面两个协程执行完毕，进而没有打印结果
	time.Sleep(5 * time.Second)
}

func responseSize(url string) {
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
	fmt.Println(len(body))
}
