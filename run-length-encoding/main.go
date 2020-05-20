package main

import (
	"fmt"
	"strconv"
)

func main() {
	input1 := "ajith"
	input2 := "ammaa"
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
}

func solve(s string) string {
	// Write your code here
	if s == "" {
		return ""
	}
	result := ""
	charS := s[0]
	countS := 1
	for i := 1; i < len(s); i++ {
		if charS == s[i] {
			countS++
		} else {
			result += strconv.Itoa(countS) + string(charS)
			charS, countS = s[i], 1
		}
	}
	return result + strconv.Itoa(countS) + string(charS)
}
