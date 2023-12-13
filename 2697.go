package main

func makeSmallestPalindrome(s string) string {
	l, r := 0, len(s)-1
	t := []byte(s)
	for l < r {
		if s[l] != s[r] {
			m := min(t[l], t[r])
			t[l] = m
			t[r] = m
		}
		l++
		r--
	}
	return string(t)
}

func min(i, j byte) byte {
	if i < j {
		return i
	}
	return j
}
