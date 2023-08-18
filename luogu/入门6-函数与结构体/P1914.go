package main

import (
	"fmt"
)

func main() {
	var n int
	var str string
	for {
		_, err := fmt.Scan(&n, &str)
		if err != nil {
			break
		}
		n %= 26
		runes := []rune(str)
		for i := 0; i < len(runes); i++ {
			c, a := int(runes[i]), int('a')
			runes[i] = rune((c-a+n)%26 + a)
		}
		fmt.Println(string(runes))
	}
}
