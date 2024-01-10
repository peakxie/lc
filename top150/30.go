package top150

// TODO 修正了很多次才修正对
func findSubstring(s string, words []string) []int { // W: func `findSubstring` is unused (unused)
	sn := len(s)
	m := len(words)
	n := len(words[0])

	res := []int{}

	for i := 0; i < n && i+m*n <= sn; i++ {
		diff := map[string]int{}

		for j := 0; j < m; j++ {
			v := s[i+j*n : i+(j+1)*n]
			diff[v]++
		}

		for _, v := range words {
			diff[v]--
			if diff[v] == 0 {
				delete(diff, v)
			}
		}

		for ii := i; ii < sn-m*n+1; ii += n {
			if ii != i {
				v := s[ii+(m-1)*n : ii+m*n]
				diff[v]++
				if diff[v] == 0 {
					delete(diff, v)
				}
				v = s[ii-n : ii]
				diff[v]--
				if diff[v] == 0 {
					delete(diff, v)
				}
			}

			if len(diff) == 0 {
				res = append(res, ii)
			}
		}
	}
	return res
}
