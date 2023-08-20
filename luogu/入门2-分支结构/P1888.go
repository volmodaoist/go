package main

import (
	"fmt"
)

// Go不支持三目运算符，非常蛋疼
func max(vals ...int) int {
	var maxv int = vals[0]
	for i := 1; i < len(vals); i++ {
		if vals[i] > maxv {
			maxv = vals[i]
		}
	}
	return maxv
}

func min(vals ...int) int {
	var minv int = vals[0]
	for i := 0; i < len(vals); i++ {
		if vals[i] < minv {
			minv = vals[i]
		}
	}
	return minv
}

func gcd(a int, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 更新之后的 min、max 使用的变长参数
func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	minv := min(a, b, c)
	maxv := max(a, b, c)
	g := gcd(minv, maxv)

	fmt.Printf("%d/%d\n", minv/g, maxv/g)
}
