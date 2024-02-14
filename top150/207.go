package top150

func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses)
	vaild := true
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}

	var dfs func(int)
	dfs = func(c int) {
		visited[c] = 1
		for _, v := range edges[c] {
			if visited[v] == 0 {
				dfs(v)
			} else if visited[v] == 1 {
				vaild = false
			}
			if !vaild {
				return
			}
		}
		visited[c] = 2
	}

	for k, _ := range edges {
		if visited[k] == 0 && vaild {
			dfs(k)
		}
	}

	return vaild
}
