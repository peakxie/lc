package top150

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	n := len(capital)

	type Pair struct {
		c int
		p int
	}

	attr := []Pair{}
	for k, v := range capital {
		attr = append(attr, Pair{c: v, p: profits[k]})
	}
	sort.Slice(attr, func(i, j int) bool { return attr[i].c < attr[j].c })

	h := &hp{}
	heap.Init(h)

	res := w
	c := 0
	for i := 0; i < k; i++ {
		for ; c < n && res >= attr[c].c; c++ {
			heap.Push(h, attr[c].p)
		}
		if h.Len() == 0 {
			break
		}
		res += heap.Pop(h).(int)
	}

	return res
}

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	t := *h
	n := len(t.IntSlice)
	x := t.IntSlice[n-1]
	h.IntSlice = t.IntSlice[:n-1]
	return x
}
