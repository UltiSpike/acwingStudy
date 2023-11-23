package main

func maxArrayValue(nums []int) int64 {
	var ans = nums[len(nums)-1]
	var sumA = nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		if sumA >= nums[i-1] {
			sumA += nums[i-1]
		} else {
			sumA = nums[i-1]
		}
		if sumA > ans {
			ans = sumA
		}
	}

	return int64(ans)
}
