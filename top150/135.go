package top150

func candy(ratings []int) int {
	n := len(ratings)

	l := make([]int, n)
	l[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
	}

	r := 1
	sum := max(l[n-1], r)
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			r++
		} else {
			r = 1
		}
		sum += max(l[i], r)
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
