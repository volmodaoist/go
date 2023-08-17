package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	arr := []int64{}
	for n != 1 {
		arr = append(arr, n)
		if n&1 == 1 {
			n = 3*n + 1
		} else {
			n /= 2
		}
	}
	arr = append(arr, int64(1))
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
