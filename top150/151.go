package top150

import "strings"

func reverseWords(s string) string {

	sl := strings.Split(s, " ")
	rl := []string{}

	for i := len(sl) - 1; i >= 0; i-- {
		v := sl[i]
		if len(v) != 0 {
			rl = append(rl, v)
		}
	}
	return strings.Join(rl, " ")
}
