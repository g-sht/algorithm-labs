package lab1

import (
	"fmt"
)

type Record struct {
	name   string
	values []int
}

func FourteenMain() {
	var n, k int
	fmt.Scan(&n, &k)

	priority := make([]int, k)
	for i := range priority {
		fmt.Scan(&priority[i])
	}

	var records []Record
	for i := 0; i < n; i++ {
		var name string
		fmt.Scan(&name)
		values := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Scan(&values[j])
		}
		records = append(records, Record{name, values})
	}

	res := mergeSortRecords(records, func(a, b Record) bool {
		for _, idx := range priority {
			vi := a.values[idx-1]
			vj := b.values[idx-1]
			if vi > vj {
				return true
			}
			if vi < vj {
				return false
			}
		}
		return false
	})

	for _, r := range res {
		fmt.Println(r.name)
	}
}

func mergeRecords(first, second []Record, comp func(i, j Record) bool) []Record {
	res := make([]Record, 0, len(first)+len(second))
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if comp(first[i], second[j]) {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, second[j])
			j++
		}
	}
	res = append(res, first[i:]...)
	res = append(res, second[j:]...)
	return res
}

func mergeSortRecords(arr []Record, comp func(i, j Record) bool) []Record {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSortRecords(arr[:mid], comp)
	right := mergeSortRecords(arr[mid:], comp)

	return mergeRecords(left, right, comp)
}
