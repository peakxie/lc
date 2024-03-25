package top150

func insert(intervals [][]int, newInterval []int) [][]int {
	s, e := newInterval[0], newInterval[1]

	res := [][]int{}
	merge := false
	for _, v := range intervals {

		if v[0] > e {
			if !merge {
				res = append(res, []int{s, e})
				merge = true
			}
			res = append(res, v)
		} else if v[1] < s {
			res = append(res, v)
		} else {
			s = min(s, v[0])
			e = max(e, v[1])
		}
	}

	if !merge {
		res = append(res, []int{s, e})
	}

	return res
}
