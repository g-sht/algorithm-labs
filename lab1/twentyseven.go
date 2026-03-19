package lab1

import (
	"fmt"
	"strconv"
)

func TwentysevenMain() {
	var n int
	fmt.Scan(&n)

	stack := make([]int, 0, n)

	for range n {
		var input string
		fmt.Scan(&input)

		if input != "+" && input != "-" {
			num, _ := strconv.Atoi(input)
			stack = append(stack, num)
			continue
		}

		el1, el2 := stack[len(stack)-2], stack[len(stack)-1]

		switch input {
		case "+":
			res := el1 + el2
			stack[len(stack)-2] = res
		case "-":
			res := el1 - el2
			stack[len(stack)-2] = res
		}

		stack = stack[:len(stack)-1]
	}

	fmt.Println(stack[0])
}
