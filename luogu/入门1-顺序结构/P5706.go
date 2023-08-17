package main

import "fmt"

func main(){
	var t float32
	var n int32
	fmt.Scan(&t)
	fmt.Scan(&n)
	fmt.Printf("%.3f\n%d\n", t / float32(n), n*2)
}