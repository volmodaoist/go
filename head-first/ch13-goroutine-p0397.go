package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// main 函数也是一个协程
func main() {
	sizes := make(chan int)
	go responseSize("https://example.com/", sizes)
	go responseSize("https://go.dev/learn/", sizes)
	// 不过访问网站，获取大小之后，打印size这种操作的先后顺序是不确定的！
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	/**
	如果不小心多写了一句，将会发现程序进入死锁，timeout之后退出
	因为 <-sizes 作为接收方会一直等待发送方的内容，如果没有发送发向通道写入。
	程序将会一直等待直至超时
	*/
}

func responseSize(url string, channel chan int) {
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
	channel <- len(body)
}
