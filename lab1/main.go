package main

import (
	"fmt"
)

func main() {
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
