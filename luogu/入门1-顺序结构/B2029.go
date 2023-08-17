package main

import (
	"fmt"
	"math"
)

func main(){
	var h, r float64
	fmt.Scan(&h, &r)
	s := h * r * r * 3.14 / 1000
	ans := math.Ceil(20.0 / s)
	fmt.Printf("%d\n", int(ans))
}