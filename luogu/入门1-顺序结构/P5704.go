package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string
	fmt.Scan(&str)
	fmt.Printf("%s\n", strings.ToUpper(str))
}