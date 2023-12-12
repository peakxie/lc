package main

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	res := []int{}
	for i := 0; i < n; i++ {
		first := true
		j := i + 1
		for ; j < n; j++ {
			if nums[i] < nums[j] {
				if first {
					first = false
				} else {
					res = append(res, nums[j])
					break
				}
			}
		}
		if j == n {
			res = append(res, -1)
		}
	}

	return res
}
