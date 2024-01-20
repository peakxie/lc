package top150

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	// 0-未扫描 1-已扫描 2-已剔除
	status := make([]int, numCourses)
	res := []int{}
	loop := false

	var dfs func(int)
	dfs = func(c int) {
		status[c] = 1
		for _, v := range graph[c] {
			if status[v] == 0 {
				dfs(v)
				if loop {
					return
				}
			} else if status[v] == 1 {
				loop = true
				return
			}
		}
		status[c] = 2
		res = append(res, c)
	}

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
	}

	for k, v := range status {
		if v == 0 {
			dfs(k)
		}
	}
	if loop {
		return []int{}
	}
	resv(res)
	return res
}

func resv(l []int) {
	n := len(l)
	for i := 0; i < n/2; i++ {
		l[i], l[n-1-i] = l[n-1-i], l[i]
	}
}
