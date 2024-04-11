package top150

func search(nums []int, target int) int {

	n := len(nums)

	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		midV := nums[mid]
		if midV == target {
			return mid
		}

		if midV >= nums[0] {
			if nums[0] <= target && target < midV {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if midV < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
