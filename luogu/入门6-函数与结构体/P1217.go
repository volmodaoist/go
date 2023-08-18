package main

import "fmt"

func isPar(x int) bool {
	cx, y := x, 0
	for cx != 0 {
		y = y*10 + (cx % 10)
		cx /= 10
	}
	return x == y
}

func isPrime(x int) bool {
	var ok bool = (x >= 2)
	for i := 2; ok && i*i <= x; i++ {
		if x%i == 0 {
			ok = false
		}
	}
	return ok
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a%2 == 0 {
		if a == 2 {
			fmt.Println(2)
		}
		a++
	}

	b = min(9999999, b) // 具有偶数位的回文数(必然包含因数11)，除了11均不是质数

	for i := a; i <= b; i += 2 {
		if isPar(i) && isPrime(i) {
			fmt.Println(i)
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
