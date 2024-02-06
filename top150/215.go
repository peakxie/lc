package top150

import (
	"container/heap"
	"math"
)

type hp []int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i] > h[j] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() any {
	old := *h
	n := len(old)
	c := old[n-1]
	*h = old[:n-1]
	return c
}

func findKthLargest(nums []int, k int) int {

	h := &hp{}
	for _, v := range nums {
		heap.Push(h, v)
	}

	res := math.MinInt
	for i := 0; i < k; i++ {
		res = heap.Pop(h).(int)
	}
	return res
}
