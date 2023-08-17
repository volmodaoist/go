package main

import (
	"fmt"
)

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	total := a * 10 + b
	fmt.Printf("%d\n", total / 19) 
}