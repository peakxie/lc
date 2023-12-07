package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func minReorder(n int, connections [][]int) int {
	lll := make([][][]int, n)
	for _, v := range connections {
		lll[v[0]] = append(lll[v[0]], []int{v[1], 1})
		lll[v[1]] = append(lll[v[1]], []int{v[0], 0})
	}
	count := 0
	var dfs func(int, int)
	dfs = func(start, p int) {
		//if start == p {
		//	return
		//}
		for _, v := range lll[start] {
			if v[0] == p { // ？
				continue
			}
			count += v[1]
			dfs(v[0], start)
		}
	}
	dfs(0, -1)
	return count
}
