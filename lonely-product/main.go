package main

import "fmt"

func main() {
	fmt.Println(solve(3, 4, 5))
	fmt.Println(solve(3, 3, 5))
	fmt.Println(solve(3, 3, 3))
}
func solve(a int, b int, c int) int {
	// Write your code here
	switch {
	case a == b && b == c:
		return 1
	case a == b:
		return c
	case b == c:
		return a
	case a == c:
		return b
	default:
		return a * b * c
	}
}
