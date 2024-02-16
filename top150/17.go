package top150

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return nil
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	res := []string{}
	var dfs func(int, string)
	dfs = func(ind int, pre string) {
		if ind == len(digits) {
			res = append(res, pre)
			return
		}
		num := digits[ind]
		for _, v := range m[num] {
			dfs(ind+1, pre+string(v))
		}
	}

	dfs(0, "")
	return res
}
