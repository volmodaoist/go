package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// 直接使用接口写入
func test1() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
}

// 间接调用接口写入
func test2() {
	var c ByteCounter
	fmt.Fprintf(&c, "hello,Dolley")
	fmt.Println(c)
}

// 通过我们是对一个文件描述符进行写入操作的走走走走                           
func main() {
	test1()
	test2()
}
