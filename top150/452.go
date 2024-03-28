package top150

import "sort"

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	base := points[0][1]
	cnt := 1
	for _, v := range points {
		if v[0] > base {
			base = v[1]
			cnt++
		}
	}
	return cnt
}
