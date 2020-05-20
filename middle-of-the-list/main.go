package main

import "fmt"

func main() {
	input1 := []int{10, 20, 30}
	input2 := []int{10, 20, 30, 40}
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))

}

func solve(nums []int) int {
	// Write your code here
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		mid--
	}
	return nums[mid]
}
