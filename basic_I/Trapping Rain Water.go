package main

import "fmt"

func trap(height []int) int {
	var nums [][2]int
	var p [2]int
	ans := 0
	for i, v := range height {
		for len(nums) > 0 && v > nums[len(nums)-1][1] {
			p = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			if len(nums) > 0 {
				ans += (min(v, nums[0][1]) - p[1]) * (i - p[0])
			}
			fmt.Println(v, p, ans)
		}
		if len(nums) == 0 || (len(nums) > 0 && v > nums[len(nums)-1][1]) {
			inline := [2]int{i, v}
			nums = append(nums, inline)
		}
	}

	return ans
}
