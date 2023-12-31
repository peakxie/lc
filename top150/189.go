package top150

func rotate(nums []int, k int) {
	n := len(nums)

	// left
	if k >= n {
		k = k % n
	}
	kk := n - k

	rev(nums, 0, kk-1)
	rev(nums, kk, n-1)
	rev(nums, 0, n-1)
}

func rev(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
