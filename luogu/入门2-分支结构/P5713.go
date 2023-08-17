package main

import "fmt"

func main()  {
	var x int
	fmt.Scan(&x)
	local := 5 * x
	luogu := 3 * x + 11
	if local < luogu {
		fmt.Println("Local")
	}else{
		fmt.Println("Luogu")
	}
}