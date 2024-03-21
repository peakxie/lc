package top150

func minWindow(s string, t string) string {

	l1, l2 := len(s), len(t)

	ori, compare := make(map[byte]int), make(map[byte]int)
	for i := 0; i < l2; i++ {
		ori[t[i]]++
	}

	check := func() bool {
		for k, v := range ori {
			if compare[k] < v {
				return false
			}
		}
		return true
	}

	l := 0
	minL := -1
	minV := l1
	for i := 0; i < l1; i++ {
		_, ok := ori[s[i]]
		if ok {
			compare[s[i]]++
		}

		for check() && l <= i {
			if minV >= i-l+1 {
				minV = i - l + 1
				minL = l
			}
			compare[s[l]]--
			l++
		}
	}
	if minL < 0 {
		return ""
	}
	return s[minL : minL+minV]
}
