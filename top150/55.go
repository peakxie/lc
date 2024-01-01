package top150

func canJump(nums []int) bool {

	n := len(nums)

	maxV := nums[0]

	for i := 0; i < n; i++ {
		if i <= maxV && nums[i]+i > maxV {
			maxV = nums[i] + i
		}

		if maxV+1 >= n {
			return true
		}
	}

	return false
}
