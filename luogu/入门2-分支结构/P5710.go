package main

import "fmt"

func checkAndPrint(value bool) {
    if value {
        fmt.Printf("1 ")
    } else {
        fmt.Printf("0 ")
    }
}

func main() {
    var x int
    fmt.Scan(&x)
    isEven := x%2 == 0
    isInRange := x > 4 && x <= 12
    
    checkAndPrint(isEven && isInRange)
    checkAndPrint(isEven || isInRange)
    checkAndPrint((isEven && !isInRange) || (!isEven && isInRange))
    checkAndPrint(!isEven && !isInRange)
    fmt.Println()
}