package top150

func generateParenthesis(n int) []string {

	ans := []string{}

	var dfs func(int, int, string)
	dfs = func(l, r int, val string) {

		if l == 0 && r == 0 {
			ans = append(ans, val)
			return
		}

		if l > 0 {
			val += "("
			dfs(l-1, r, val)
			val = val[:len(val)-1]
		}

		if r > l {
			val += ")"
			dfs(l, r-1, val)
			val = val[:len(val)-1]
		}
	}

	dfs(n, n, "")
	return ans
}
