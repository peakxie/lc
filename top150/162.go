package top150

func findPeakElement(nums []int) int {

	n := len(nums)

	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}

	return n - 1
}
