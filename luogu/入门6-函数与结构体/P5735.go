package main

import (
	"fmt"
	"math"
)

func calc(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	p1 := math.Pow(x1-x2, 2)
	p2 := math.Pow(y1-y2, 2)
	return math.Sqrt(p1 + p2)
}

func main() {
	var coordinates [3][2]float64

	// Go少加了\n导致读入出错，luogo上面 WA 好几次
	for i := 0; i < 3; i++ {
		fmt.Scanf("%f %f\n", &coordinates[i][0], &coordinates[i][1])
	}

	ans := 0.0
	for i := 0; i < 3; i++ {
		ans += calc(coordinates[i][0], coordinates[i][1],
			coordinates[(i+1)%3][0], coordinates[(i+1)%3][1])
	}
	fmt.Printf("%.2f\n", ans)
}
