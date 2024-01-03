package top150

import "sort"

func hIndex(citations []int) int {

	n := len(citations)
	sort.Ints(citations)
	h := 0
	for i := n - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}

	return h
}
