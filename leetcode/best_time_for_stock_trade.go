package main

import (
	"math"
)

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]

	if prices == nil || len(prices) <= 1 {
		return res
	}

	for i := 0; i < len(prices); i++ {
		if prices[i] > min {
			res = int(math.Max(float64(res), float64(prices[i]-min)))
		} else {
			min = prices[i]
		}
	}

	return res
}

// func main() {
// 	prices := []int{2, 4, 1}

// 	fmt.Println(maxProfit(prices))

// }
