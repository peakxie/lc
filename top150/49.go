package top150

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}

	res := [][]string{}

	for _, str := range strs {
		l := [26]int{}
		for _, ch := range str {
			v := ch - 'a'
			l[v]++
		}
		m[l] = append(m[l], str)
	}
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
