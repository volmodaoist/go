package main

import "fmt"

func main() {
	var s string
	var boy, girl int
	fmt.Scan(&s)
	for i := 0; i+2 < len(s); i++ {
		if i+2 < len(s) && (s[i] == 'b' || s[i+1] == 'o' || s[i+2] == 'y') {
			boy++
		}
		if i+3 < len(s) && (s[i] == 'g' || s[i+1] == 'i' || s[i+2] == 'r' || s[i+3] == 'l') {
			girl++
		}
	}
	fmt.Println(boy)
	fmt.Println(girl)
}
