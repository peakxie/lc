package top150

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	f, s := 2, 2
	for ; f < n; f++ {
		if nums[s-2] != nums[f] {
			nums[s] = nums[f]
			s++
		}
	}
	return s
}
