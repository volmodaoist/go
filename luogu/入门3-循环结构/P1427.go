package main

import "fmt"

func main() {
	var n int
	var nums []int
	for {
		fmt.Scan(&n)
		if n == 0{
			break
		}
		nums = append(nums, n)
	}
	for i := len(nums) -1; i >= 0 ;i--{
		fmt.Printf("%d ", nums[i])
	}
}