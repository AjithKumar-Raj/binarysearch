package main

import "fmt"

func main() {
	input1 := []int{5, 6}
	input2 := []int{5}
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
}

func solve(nums []int) bool {
	// Write your code here
	if len(nums) == 1 {
		return true
	}
	return false
}
