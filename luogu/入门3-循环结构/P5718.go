package main

import "fmt"

func main() {
	var n int
	var minv int = 2e9
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {

		fmt.Scan(&arr[i])
		if minv > arr[i] {
			minv = arr[i]
		}
	}
	fmt.Println(minv)
}
