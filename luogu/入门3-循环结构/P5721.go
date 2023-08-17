package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Printf("%02d", cnt)
			cnt++
		}
		fmt.Println()
	}
}
