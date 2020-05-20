package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solve([]int{2, 3, 5}))
}

func solve(nums []int) bool {
	// Write your code here
	sort.Ints(nums)
	return nums[len(nums)-1] > (nums[len(nums)-2] * 2)

}
