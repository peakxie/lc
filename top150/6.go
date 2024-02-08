package top150

func convert(s string, numRows int) string {
	n := len(s)
	r := numRows
	if r == 1 || r >= n {
		return s
	}

	t := r*2 - 2
	c := (n+t-1)/t*r - 1

	mat := make([][]byte, r)
	for i := 0; i < r; i++ {
		mat[i] = make([]byte, c)
	}

	x, y := 0, 0
	for i := 0; i < n; i++ {
		mat[x][y] = s[i]
		if i%t < r-1 {
			x++
		} else {
			x--
			y++
		}
	}

	res := []byte{}
	for _, l := range mat {
		for _, v := range l {
			if v > 0 {
				res = append(res, v)
			}
		}
	}

	return string(res)
}
