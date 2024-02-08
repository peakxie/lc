package top150

type tb struct {
	val    int
	symbal string
}

func intToRoman(num int) string {
	l := []tb{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var res string
	for _, v := range l {
		for num >= v.val {
			num -= v.val
			res += v.symbal
		}
	}

	return res
}
