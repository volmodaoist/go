package main

import "fmt"

func prod(n int64, arr []int64) int64 {
	if arr[n] != 0 {
		return arr[n]
	}
	arr[n] = n * prod(n-1, arr)
	return arr[n]
}

// 下面这个其实没有办法AC，需要使用高精度算法，
// 不过我们仅仅是为了熟悉语法而已，因而也就不追求全对了
func main() {
	var n, s int64
	fmt.Scan(&n)
	arr := make([]int64, n+1)
	arr[0], arr[1] = 1, 1

	for i := int64(1); i <= n; i++ {
		s += prod(i, arr)
	}

	fmt.Println(s)
}
