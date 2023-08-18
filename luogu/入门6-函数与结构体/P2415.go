package main

import "fmt"

func main() {
	var ans int64 = 0
	var n, t int
	for {
		_, err := fmt.Scanf("%d", &t)
		if err != nil {
			break
		}
		ans += int64(t)
		n++
	}
	ans = ans << (n - 1)
	fmt.Printf("%d\n", ans)
}
