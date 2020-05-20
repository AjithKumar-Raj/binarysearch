package main

import (
	"fmt"
	"strings"
)

func main() {
	input1 := ""
	input2 := "hippo"
	input3 := "hippollama"
	fmt.Println(solve(input1))
	fmt.Println(solve(input2))
	fmt.Println(solve(input3))

}

func solve(s string) bool {
	// Write your code here
	return strings.Count(s, "llama") == strings.Count(s, "hippo")
}
