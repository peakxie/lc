package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	ll := make([][]int, len(roads)+1)
	for _, v := range roads {
		ll[v[0]] = append(ll[v[0]], v[1])
		ll[v[1]] = append(ll[v[1]], v[0])
	}

	var res int64
	var dfs func(int, int) int
	dfs = func(cur, fa int) int {
		num := 1
		for _, vv := range ll[cur] {
			if vv != fa {
				cnt := dfs(vv, cur)
				num = num + cnt
				res = res + int64((cnt+seats-1)/seats)
			}

		}
		return num
	}
	dfs(0, -1)
	return res
}
