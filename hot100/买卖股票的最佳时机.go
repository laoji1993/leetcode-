package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	var maxProfit int
	for _,v := range prices {
		if v < minPrice {
			minPrice = v
		} else if v - minPrice > maxProfit {
			maxProfit = v -minPrice
		}
	}
	return maxProfit
}
