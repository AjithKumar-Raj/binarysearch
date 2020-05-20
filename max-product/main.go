package main

import (
	"fmt"
	"sort"
)

func main() {
	input1 := []int{}
	input2 := []int{}
	fmt.Println(input1)
	fmt.Println(input2)
}

func solve(nums []int) int {
	// Write your code here
	sort.Ints(nums)
	ln := len(nums)
	a := nums[0] * nums[1]
	b := nums[ln-1] * nums[ln-2]
	if a < b {
		return b
	}
	return a
}
