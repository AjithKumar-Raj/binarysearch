package main

import (
	"fmt"
	"math"
)

func main() {
	input1 := []float32{1.43, 4.23}
	fmt.Println(solve(input1))
}

func solve(prices []float32) []float32 {
	// Write your code here
	for i := range prices {
		prices[i] = float32(math.Round(float64(prices[i]*1.15*100)) / 100)
	}
	return prices
}
