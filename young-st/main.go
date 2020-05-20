package main

import "fmt"

func main() {
	input1 := []int{8, 5, 9}
	input2 := []int{1, 3, 7}
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
}

func solve(nums []int) bool {
	// Write your code here
	for _, num := range nums {
		if num == 1 || num == 3 || num == 7 {
			return true
		}
	}
	return false
}
