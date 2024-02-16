package top150

func minMutation(startGene string, endGene string, bank []string) int {
	res := 0
	if startGene == endGene {
		return res
	}

	m := map[string]bool{}
	for _, v := range bank {
		m[v] = true
	}

	if !m[endGene] {
		return -1
	}

	q := []string{startGene}
	for len(q) != 0 {
		t := q
		q = nil
		res++
		for _, v := range t {
			for i, ch := range v {
				for _, x := range "ACGT" {
					if ch != x {
						gene := v[:i] + string(x) + v[i+1:]
						if _, ok := m[gene]; ok {
							if gene == endGene {
								return res
							}
							delete(m, gene)
							q = append(q, gene)
						}
					}
				}
			}
		}
	}
	return -1
}
