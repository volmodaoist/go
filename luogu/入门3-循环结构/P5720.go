package main

import "fmt"

func main() {
	var x, cnt int
	fmt.Scanf("%d", &x)
	for x > 1 {
		x /= 2
		cnt++
	}
	fmt.Println(cnt + 1)
}
