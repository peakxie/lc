package top150

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	m := map[byte]int{}

	e := 0
	res := 0
	for i := 0; i < n; i++ {
		for ; e < n && m[s[e]] == 0; e++ {
			m[s[e]]++
			res = max(res, e-i+1)
		}
		delete(m, s[i])
	}

	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
