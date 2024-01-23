package top150

func calculate(s string) int {

	ops := []int{1}
	sign := 1

	res := 0
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && int(s[i]-'0') >= 0 && int(s[i]-'0') <= 9; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += sign * num
		}

	}

	return res
}
