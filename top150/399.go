package top150

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	m := map[string]int{}
	for _, v := range equations {
		for i := 0; i < len(v); i++ {
			_, ok := m[v[i]]
			if !ok {
				m[v[i]] = len(m)
			}
		}
	}

	type Node struct {
		to  int
		val float64
	}

	graph := make([][]Node, len(m))
	for k, v := range equations {
		graph[m[v[0]]] = append(graph[m[v[0]]], Node{to: m[v[1]], val: values[k]})
		graph[m[v[1]]] = append(graph[m[v[1]]], Node{to: m[v[0]], val: 1 / values[k]})
	}

	bfs := func(s, e int) float64 {
		vals := make([]float64, len(m))
		vals[s] = 1
		queue := []int{s}
		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]

			if v == e {
				return vals[v]
			}
			for _, vv := range graph[v] {
				if vals[vv.to] == 0 {
					vals[vv.to] = vals[v] * vv.val
					queue = append(queue, vv.to)
				}
			}
		}
		return -1
	}

	res := []float64{}
	for _, v := range queries {
		id0, ok0 := m[v[0]]
		id1, ok1 := m[v[1]]
		if !ok0 || !ok1 {
			res = append(res, -1)
			continue
		}
		res = append(res, bfs(id0, id1))
	}
	return res
}
