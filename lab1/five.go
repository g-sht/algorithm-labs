package main

import (
	"fmt"
	"strconv"
)

func fiveMain() {
	var firstNum string
	var secondNum string
	var operand string

	fmt.Scan(&firstNum, &operand, &secondNum)

	isFirstNumNeg := false
	if firstNum[0] == '-' {
		isFirstNumNeg = true
		firstNum = firstNum[1:]
	}

	isSecondNumNeg := false
	if secondNum[0] == '-' {
		isSecondNumNeg = true
		secondNum = secondNum[1:]
	}

	firstNumArr := strToRevIntSlice(firstNum)
	secondNumArr := strToRevIntSlice(secondNum)

	var result []int
	var resultIsNegative bool

	switch {
	case !isFirstNumNeg && !isSecondNumNeg && operand == "+": // A + B
		result = longOperations(firstNumArr, secondNumArr, "+")
		resultIsNegative = false
	case !isFirstNumNeg && isSecondNumNeg && operand == "+": // A + (-B) = A - B
		if compareAbs(firstNumArr, secondNumArr) >= 0 {
			result = longOperations(firstNumArr, secondNumArr, "-")
			resultIsNegative = false
		} else {
			result = longOperations(secondNumArr, firstNumArr, "-")
			resultIsNegative = true
		}
	case isFirstNumNeg && !isSecondNumNeg && operand == "+": // (-A) + B = B - A
		if compareAbs(secondNumArr, firstNumArr) >= 0 {
			result = longOperations(secondNumArr, firstNumArr, "-")
			resultIsNegative = false
		} else {
			result = longOperations(firstNumArr, secondNumArr, "-")
			resultIsNegative = true
		}
	case isFirstNumNeg && isSecondNumNeg && operand == "+": // (-A) + (-B) = -(A + B)
		result = longOperations(firstNumArr, secondNumArr, "+")
		resultIsNegative = true

	case !isFirstNumNeg && !isSecondNumNeg && operand == "-": // A - B
		if compareAbs(firstNumArr, secondNumArr) >= 0 {
			result = longOperations(firstNumArr, secondNumArr, "-")
			resultIsNegative = false
		} else {
			result = longOperations(secondNumArr, firstNumArr, "-")
			resultIsNegative = true
		}
	case !isFirstNumNeg && isSecondNumNeg && operand == "-": // A - (-B) = A + B
		result = longOperations(firstNumArr, secondNumArr, "+")
		resultIsNegative = false
	case isFirstNumNeg && !isSecondNumNeg && operand == "-": // (-A) - B = -(A + B)
		result = longOperations(firstNumArr, secondNumArr, "+")
		resultIsNegative = true
	case isFirstNumNeg && isSecondNumNeg && operand == "-": // (-A) - (-B) = B - A
		if compareAbs(secondNumArr, firstNumArr) >= 0 {
			result = longOperations(secondNumArr, firstNumArr, "-")
			resultIsNegative = false
		} else {
			result = longOperations(firstNumArr, secondNumArr, "-")
			resultIsNegative = true
		}
	}

	if resultIsNegative {
		fmt.Print("-")
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
	fmt.Println()
}

func longOperations(num1, num2 []int, operand string) []int {
	out := make([]int, 0, max(len(num1), len(num2)))
	maxSize := max(len(num1), len(num2))

	switch operand {
	case "+":
		carry := 0
		for i := 0; i < maxSize || carry > 0; i++ {
			sum := carry
			if i < len(num1) {
				sum += num1[i]
			}
			if i < len(num2) {
				sum += num2[i]
			}
			out = append(out, sum%10)
			carry = sum / 10
		}
	case "-":
		borrow := 0
		for i := 0; i < len(num1); i++ {
			diff := num1[i] - borrow
			if i < len(num2) {
				diff -= num2[i]
			}

			if diff < 0 {
				diff += 10
				borrow = 1
			} else {
				borrow = 0
			}
			out = append(out, diff)
		}

		for len(out) > 1 && out[len(out)-1] == 0 {
			out = out[:len(out)-1]
		}
	}

	return out
}

func compareAbs(a, b []int) int {
	if len(a) > len(b) {
		return 1
	}
	if len(a) < len(b) {
		return -1
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		}
		if a[i] < b[i] {
			return -1
		}
	}
	return 0
}

func strToRevIntSlice(s string) []int {
	res := make([]int, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(s[i]))
		res = append(res, digit)
	}
	return res
}
