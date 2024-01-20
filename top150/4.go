package top150

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	k := m + n
	if k%2 == 1 {
		return float64(getMid(nums1, nums2, k/2+1))
	}

	return float64(getMid(nums1, nums2, k/2)+getMid(nums1, nums2, k/2+1)) / 2
}

func getMid(nums1 []int, nums2 []int, k int) int {
	ind1, ind2 := 0, 0
	for {
		if ind1 > len(nums1)-1 {
			return nums2[ind2+k-1]
		}
		if ind2 > len(nums2)-1 {
			return nums1[ind1+k-1]
		}
		if k == 1 {
			return min(nums1[ind1], nums2[ind2])
		}

		half := k / 2
		newInd1 := min(ind1+half, len(nums1)) - 1
		newInd2 := min(ind2+half, len(nums2)) - 1
		v1 := nums1[newInd1]
		v2 := nums2[newInd2]
		if v1 >= v2 {
			k -= (newInd2 + 1 - ind2)
			ind2 = newInd2 + 1
		} else {
			k -= (newInd1 + 1 - ind1)
			ind1 = newInd1 + 1
		}
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
