package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&arr[i])
		ans := 0
		for j := 1; j < i; j++ {
			if arr[j] < arr[i] {
				ans++
			}
		}
		fmt.Printf("%d ", ans)
	}
	fmt.Println()
}
