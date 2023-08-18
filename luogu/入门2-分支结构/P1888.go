package main

import (
	"fmt"
)

// Go不支持三目运算符，非常蛋疼
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func gcd(a int, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	minv := min(a, min(b, c))
	maxv := max(a, max(b, c))
	g := gcd(minv, maxv)

	fmt.Printf("%d/%d\n", minv/g, maxv/g)
}
