package top150

import "math"

func minSubArrayLen(target int, nums []int) int {

	n := len(nums)

	sum := 0
	s, e := 0, 0
	res := math.MaxInt
	for e < n {
		sum += nums[e]
		for s <= e && sum >= target { // s <= e也不需要判断，因为过了一定是负值
			//if sum == target { 和大于等目标值的连续数组长度最小值
			res = min(res, e-s+1)
			//}
			sum -= nums[s]
			s++
		}
		e++
	}
	if res == math.MaxInt {
		res = 0
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
