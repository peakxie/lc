package top150

func spiralOrder(matrix [][]int) []int {
	u, d := 0, len(matrix)-1
	l, r := 0, len(matrix[0])-1

	res := []int{}

	for {
		i, j := u, l
		for ; j <= r; j++ {
			res = append(res, matrix[i][j])
		}
		u++
		if u > d {
			break
		}
		i = u
		j = r
		for ; i <= d; i++ {
			res = append(res, matrix[i][j])
		}
		r--
		if l > r {
			break
		}
		i = d
		j = r
		for ; j >= l; j-- {
			res = append(res, matrix[i][j])
		}
		d--
		if u > d {
			break
		}
		i = d
		j = l
		for ; i >= u; i-- {
			res = append(res, matrix[i][j])
		}
		l++
		if l > r {
			break
		}
	}
	return res
}
