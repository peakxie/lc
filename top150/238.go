package top150

func productExceptSelf(nums []int) []int {

	n := len(nums)

	res := make([]int, n)

	ll := make([]int, n)
	ll[0] = 1

	for i := 1; i < n; i++ {
		ll[i] = ll[i-1] * nums[i-1]
	}

	r := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = ll[i] * r

		r *= nums[i]
	}
	return res
}
