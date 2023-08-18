package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	hash := [10]int{}
	for i := m; i <= n; i++ {
		x := i
		for x != 0 {
			hash[x%10]++
			x /= 10
		}
	}
	for _, e := range hash {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}
