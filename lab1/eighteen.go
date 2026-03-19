package lab1

import (
	"fmt"
	"math"
)

func EighteenMain() {
	var n, k int
	fmt.Scan(&n, &k)

	diff := make([]int, n+1)

	for range k {
		var first, last int
		fmt.Scan(&first, &last)

		diff[first] += 1
		diff[last] -= 1
	}

	minCams := math.MaxInt32
	currentLifts := 0

	for i := 1; i < n; i++ {
		currentLifts += diff[i]
		minCams = min(minCams, currentLifts)
	}

	if minCams != math.MaxInt32 {
		fmt.Println(minCams)
		return
	}
	fmt.Println(-1)
}
