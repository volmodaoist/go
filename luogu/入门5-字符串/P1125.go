package main

import "fmt"

func isPrime(x int) bool {
	var ok bool = (x >= 2)
	for i := 2; ok && i < x; i++ {
		if x%i == 0 {
			ok = false
		}
	}
	return ok
}

func main() {
	var s string
	for {
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		hash := make(map[uint8]int)
		for i := 0; i < len(s); i++ {
			hash[s[i]]++
		}
		minv, maxv := int(2e9), int(-2e9)
		for _, cnt := range hash {
			if cnt > maxv {
				maxv = cnt
			}
			if cnt < minv {
				minv = cnt
			}
		}
		if isPrime(maxv - minv) {
			fmt.Println("Lucky Word")
			fmt.Println(maxv - minv)
		} else {
			fmt.Println("No Answer")
			fmt.Println(0)
		}
	}
}
