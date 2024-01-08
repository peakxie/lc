package top150

func trap(height []int) int {

	n := len(height)
	l := make([]int, n)

	l[0] = 0
	for i := 1; i < n; i++ {
		l[i] = max(height[i-1], l[i-1])
	}

	sum := 0
	r := height[n-1]
	for i := n - 2; i >= 0; i-- {
		r = max(r, height[i+1])
		z := min(l[i], r) - height[i]
		if z > 0 {
			sum += z
		}
	}

	return sum

}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
