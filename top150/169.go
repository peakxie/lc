package top150

import "math"

func majorityElement(nums []int) int {
	count := 0
	val := math.MaxInt

	for _, v := range nums {
		if count == 0 {
			val = v
		}
		if val == v {
			count++
		} else {
			count--
		}
	}
	return val
}
