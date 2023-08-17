package main

import "fmt"

func main() {
	var x int
	var c string
	fmt.Scan(&x)
	if x > 1 {
		c = "s"
	}
	fmt.Printf("Today, I ate %d apple%s.\n", x, c)
}
