package main

import "fmt"

func main() {
	input1 := []int{1, 2, 3, 4}
	input2 := []int{5}
	input3 := []int{1, 4, 3, 1}
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
	fmt.Println(solve(input3))
}
func solve(nums []int) bool {
	// Write your code here
	return nums[0] == nums[len(nums)-1] && len(nums) > 1
}
