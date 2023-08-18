package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	isbn, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	isbn = strings.TrimRight(isbn, "\n")

	sumv, cnt, length := 0, byte(1), len(isbn)
	for i := 0; i < length-1; i++ {
		if isbn[i] != '-' {
			sumv += int((isbn[i] - '0') * cnt)
			cnt++
		}
	}
	sumv %= 11

	if sumv == 10 && isbn[length-1] == 'X' || sumv < 10 && sumv == int(isbn[length-1]-'0') {
		fmt.Printf("Right\n")
	} else {
		isbn := []rune(isbn)
		if sumv < 10 {
			isbn[length-1] = rune('0' + sumv)
		} else {
			isbn[length-1] = rune('X')
		}
		fmt.Println(string(isbn))
	}
}
