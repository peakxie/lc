package top150

import "strings"

func simplifyPath(path string) string {
	res := []string{}

	s := strings.Split(path, "/")
	for _, v := range s {
		switch v {
		case "":
		case ".":
		case "/":
		case "..":
			resL := len(res)
			if resL != 0 {
				res = res[:resL-1]
			}
		default:
			res = append(res, v)
		}
	}

	return "/" + strings.Join(res, "/")
}
