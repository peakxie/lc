package main

import "fmt"

func main() {
	nums := []int{8, 8, 8, 8}
	fmt.Println(minD(nums))
}

func minD(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	count := len(nums)
	delNum := 0

	for i := 0; i < count; i++ {
		if (i+delNum)%2 == 0 {
			for j := i + 1; j < count; j++ {
				if nums[i] == nums[j] {
					delNum++
					i = j
					fmt.Printf("del(%d): %d\n", j, nums[j])
				} else {
					i = j
					break
				}
			}
		}
	}
	if (count-delNum)%2 != 0 {
		delNum++
	}
	return delNum
}
