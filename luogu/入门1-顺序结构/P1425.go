package main

import (
	"fmt"
)

func main(){
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	bg, ed := a * 60 + b, c * 60 + d 
	fmt.Printf("%d %d\n", (ed - bg) / 60, (ed - bg) % 60)
}