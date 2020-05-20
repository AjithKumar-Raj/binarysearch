package main

import (
	"fmt"
)

func main() {
	input1 := []int{1, 2, 3, 4, 5}
	input2 := []int{1, 2, 3, 4, 5}
	fmt.Println(solve(input1, input2))
}

func solve(a []int, b []int) int {
	// Write your code here
	result := 0
	if len(a) > 0 {
		for i := 0; i < len(a); i++ {
			result += a[i] * b[i]
		}
	}
	return result

}
