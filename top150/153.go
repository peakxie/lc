package top150

func findMin(nums []int) int {

	n := len(nums)

	l, r := 0, n-1
	res := nums[0]
	for l <= r {
		mid := (l + r) / 2
		if nums[l] <= nums[mid] {
			res = min(res, nums[l])
			l = mid + 1
		} else {
			res = min(res, nums[mid])
			r = mid - 1
		}
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
