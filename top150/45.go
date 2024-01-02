package top150

func jump(nums []int) int {
	n := len(nums)

	e := 0
	count := 0
	maxV := 0

	for i := 0; i < n-1; i++ {
		maxV = max(maxV, i+nums[i])
		if i == e {
			e = maxV
			count++
		}
	}
	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
