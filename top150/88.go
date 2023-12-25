package top150

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		cur := 0
		if i < 0 {
			cur = nums2[j]
			j--
		} else if j < 0 {
			cur = nums1[i]
			i--
		} else if nums1[i] > nums2[j] {
			cur = nums1[i]
			i--
		} else {
			cur = nums2[j]
			j--
		}

		nums1[tail] = cur
	}
}
