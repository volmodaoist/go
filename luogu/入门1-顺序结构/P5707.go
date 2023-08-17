package main

import "fmt"

func main(){
	var s, v int
	fmt.Scan(&s, &v)
	base := 8 * 60 + 24 * 60
	t := (s + v - 1) / v + 10
	base -= t
	if base >= 24 * 60{
		base -= 24 * 60
	}
	fmt.Printf("%02d:%02d\n", base / 60, base % 60)
}