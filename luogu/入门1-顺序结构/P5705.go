package main

import "fmt"

// 翻转一个字符串
func main() {
	var str string
	fmt.Scan(&str)
	runes := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Printf("%s\n", string(runes))
}
