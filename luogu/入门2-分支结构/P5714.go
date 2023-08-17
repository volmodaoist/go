package main

import (
	"fmt"
)

func main() {
	var m, h float32
	for {
		_, err := fmt.Scan(&m, &h)
		if err != nil {
			break
		}
		bmi := m / (h * h)
		if bmi < 18.5 {
			fmt.Println("Underweight")
		} else if bmi >= 18.5 && bmi < 24 {
			fmt.Println("Normal")
		} else {
			fmt.Printf("%.f\n", bmi)
			fmt.Println("Overweight")
		}
	}

}
