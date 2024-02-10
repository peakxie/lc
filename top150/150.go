package top150

import "strconv"

func evalRPN(tokens []string) int {
	res := 0
	stack := []int{}
	for _, v := range tokens {
		vv, err := strconv.Atoi(v)
		if err != nil {
			n := len(stack)
			r := 0
			n1, n2 := stack[n-2], stack[n-1]
			stack = stack[:n-2]
			switch v {
			case "+":
				r = n1 + n2
			case "-":
				r = n1 - n2
			case "*":
				r = n1 * n2
			case "/":
				r = n1 / n2
			}
			stack = append(stack, r)
		} else {
			stack = append(stack, vv)
		}
	}
	if len(stack) != 0 {
		res = stack[0]
	}
	return res
}
