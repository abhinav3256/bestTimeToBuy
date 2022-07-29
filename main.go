package main

import "fmt"

func main() {
	// test()
	arr := []int{1, 2, 4, 3}
	result := maxProfit(arr)

	fmt.Println(result)
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	result := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		next := temp + prices[i] - prices[i-1]
		result = max(result, next)
		if next >= 0 {
			temp = next
		} else {
			temp = 0
		}
		fmt.Println("next", next)
		fmt.Println("temp", temp)
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
