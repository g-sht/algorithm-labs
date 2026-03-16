package main

func simDiff(a, b []int) []int {
	simDiff := make([]int, 0, 2)

	quickSortThree(a, 0, len(a)-1)
	quickSortThree(b, 0, len(b)-1)

	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			simDiff = append(simDiff, a[i])
			i++
		} else if a[i] > b[j] {
			simDiff = append(simDiff, b[j])
			j++
		} else {
			i++
			j++
			continue
		}
	}

	if i < len(a) {
		simDiff = append(simDiff, a[i:]...)
	} else if j < len(b) {
		simDiff = append(simDiff, b[j:]...)
	}

	return simDiff
}

func quickSortThree(arr []int, start, end int) {
	if start >= end {
		return
	}

	pivot := start + (end-start)/2

	arr[end], arr[pivot] = arr[pivot], arr[end]
	swapInd := start
	for i := start; i < end; i++ {
		if arr[i] < arr[end] {
			arr[i], arr[swapInd] = arr[swapInd], arr[i]
			swapInd++
		}
	}

	arr[swapInd], arr[end] = arr[end], arr[swapInd]
	quickSortThree(arr, start, swapInd-1)
	quickSortThree(arr, swapInd+1, end)
}
