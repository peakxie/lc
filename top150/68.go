package top150

import "strings"

func fullJustify(words []string, maxWidth int) []string {

	res := []string{}

	n := len(words)
	r := 0
	for {

		l := r
		sumLen := 0
		for r < n && len(words[r])+sumLen+r-l <= maxWidth {
			sumLen += len(words[r])
			r++
		}

		// 最后一行
		if r == n {
			tmp := maxWidth - sumLen - (r - l - 1)
			res = append(res, strings.Join(words[l:r], repeatSpace(1))+repeatSpace(tmp))
			return res
		}

		nums := r - l
		// 只有一个单词行
		if nums == 1 {
			tmp := maxWidth - sumLen
			res = append(res, words[l]+repeatSpace(tmp))
			continue
		}

		space := (maxWidth - sumLen) / (nums - 1)
		extSpace := (maxWidth - sumLen) % (nums - 1)
		s1 := strings.Join(words[l:l+extSpace+1], repeatSpace(space+1))
		s2 := strings.Join(words[l+extSpace+1:r], repeatSpace(space))
		res = append(res, s1+repeatSpace(space)+s2)
	}

}

func repeatSpace(i int) string {
	return strings.Repeat(" ", i)
}
