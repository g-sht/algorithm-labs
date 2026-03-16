package main

import (
	"fmt"
)

func sevenMain() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	totalSum := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		totalSum += arr[i]
	}

	dp := make([]bool, totalSum+1)
	dp[0] = true

	for _, w := range arr {
		for j := totalSum; j >= w; j-- {
			if dp[j-w] {
				dp[j] = true
			}
		}
	}

	isMinInit := false
	minDiff := 0
	for s := range totalSum/2 + 1 {
		if dp[s] {
			diff := abs(totalSum - 2*s)
			if !isMinInit {
				minDiff = diff
				isMinInit = true
			} else {
				minDiff = min(minDiff, diff)
			}
		}
	}

	fmt.Println(minDiff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
