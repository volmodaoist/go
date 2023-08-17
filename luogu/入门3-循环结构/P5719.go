package main

import "fmt"

func main() {
	var n, k int
	var a, b, cnt_a, cnt_b int
	fmt.Scanf("%d %d", &n, &k)
	for i := 1; i <= n; i++ {
		if i%k == 0 {
			a += i
			cnt_a++
		} else {
			b += i
			cnt_b++
		}
	}
	fmt.Printf("%.1f %.1f", float32(a)/float32(cnt_a), float32(b)/float32(cnt_b))
}
