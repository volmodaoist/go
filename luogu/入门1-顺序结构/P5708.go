package main

import (
    "fmt"
    "math"
)

func calc(a float64, b float64, c float64) float64 {
    p := float64(a+b+c) / 2
    return math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))
}

func main() {
    var a, b, c float64
    fmt.Scan(&a, &b, &c)
    fmt.Printf("%.1f\n", calc(a, b, c))
}
