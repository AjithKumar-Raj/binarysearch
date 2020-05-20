package main

import "fmt"

func main() {
	input1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	input2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
}

func solve(lst []int) []int {
	// Write your code here
	res := make([]int, len(lst))
	for i := range lst {
		for j := i + 1; j < len(lst); j++ {
			if lst[j] < lst[i] {
				res[i]++
			}
		}
	}
	return res
}
