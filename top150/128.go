package top150

func longestConsecutive(nums []int) int {

	m := map[int]bool{}

	for _, v := range nums {
		m[v] = true
	}

	res := 0
	for k, _ := range m {
		if !m[k-1] {
			cur := k
			count := 1
			for m[cur+1] {
				cur++
				count++
			}
			res = max(res, count)
		}
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
