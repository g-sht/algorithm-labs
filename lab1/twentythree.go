package lab1

import (
	"bufio"
	"fmt"
	"os"
)

type Contestant struct {
	name     string
	problems int
	penalty  int
}

func NewContestant(name string, problems, penalty int) *Contestant {
	return &Contestant{name, problems, penalty}
}

func CompareContestants(a, b *Contestant) bool {
	if a.problems > b.problems {
		return true
	} else if a.problems == b.problems && a.penalty < b.penalty {
		return true
	} else if a.problems == b.problems && a.penalty == b.penalty && a.name < b.name {
		return true
	}

	return false
}

type Heap struct {
	items []*Contestant
}

func NewHeapWithHeapify(items []*Contestant) *Heap {
	h := &Heap{items}
	for i := len(items)/2 - 1; i >= 0; i-- {
		h.Heapify(len(h.items), i)
	}

	return h
}

func (h *Heap) Heapify(size, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < size && CompareContestants(h.items[largest], h.items[l]) {
		largest = l
	}

	if r < size && CompareContestants(h.items[largest], h.items[r]) {
		largest = r
	}

	if largest != i {
		h.items[i], h.items[largest] = h.items[largest], h.items[i]
		h.Heapify(size, largest)
	}
}

func (h *Heap) HeapSort() []*Contestant {
	n := len(h.items)

	for i := n - 1; i > 0; i-- {
		h.items[0], h.items[i] = h.items[i], h.items[0]

		h.Heapify(i, 0)
	}

	return h.items
}

func TwentythreeMain() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	contestants := make([]*Contestant, 0, n)
	for range n {
		var name string
		var problems, penalty int

		fmt.Fscan(in, &name, &problems, &penalty)

		contestant := NewContestant(name, problems, penalty)
		contestants = append(contestants, contestant)
	}

	heap := NewHeapWithHeapify(contestants)
	ans := heap.HeapSort()

	for _, contestant := range ans {
		fmt.Fprintln(out, contestant.name)
	}
}
