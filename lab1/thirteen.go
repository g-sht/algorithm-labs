package main

import (
	"bufio"
	"fmt"
	"os"
)

func thirteenMain() {
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(r, &n)

	var min1, min2, max1, max2, max3 int

	const INF = 10000001

	max1, max2, max3 = -INF, -INF, -INF
	min1, min2 = INF, INF

	for range n {
		var el int
		fmt.Fscan(r, &el)

		if el < min1 {
			min1, min2 = el, min1
		} else if el < min2 {
			min2 = el
		}

		if el > max1 {
			max1, max2, max3 = el, max1, max2
		} else if el > max2 {
			max2, max3 = el, max2
		} else if el > max3 {
			max3 = el
		}
	}

	var isFirstAns, isSecondAns bool
	var ans1, ans2 int64

	if max1 != INF && max2 != INF && max3 != INF {
		ans1 = int64(max1) * int64(max2) * int64(max3)
		isFirstAns = true
	}
	if min1 != -INF && min2 != INF && max1 != INF {
		ans2 = int64(min1) * int64(min2) * int64(max1)
		isSecondAns = true
	}

	if isFirstAns && isSecondAns {
		if ans1 > ans2 {
			fmt.Println(ans1)
		} else {
			fmt.Println(ans2)
		}
	} else if !isFirstAns && isSecondAns {
		fmt.Println(ans2)
	} else if isFirstAns {
		fmt.Println(ans1)
	} else {
		fmt.Println(0)
	}
}
