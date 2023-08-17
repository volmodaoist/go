package main

import "fmt"

func main(){
    var x string
    num_space := [3]int {2, 1, 0}
    num_char := [3]int  {1, 3, 5}
    fmt.Scan(&x)
	for i:= 0; i < 3; i++ {
		for j := 0; j < num_space[i]; j++ { fmt.Printf(" ") }
		for j := 0; j < num_char[i];  j++ { fmt.Printf("%s", x)}
		for j := 0; j < num_space[i]; j++ { fmt.Printf(" ") }
		fmt.Printf("\n")
	}
}