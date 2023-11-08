package main

import "fmt"

func main() {
	var n, q int
	_, _ = fmt.Scanln(&n, &q)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}

	query := make([]int, q)

	for i := 0; i < q; i++ {
		_, _ = fmt.Scanln(&query[i])

	}
	for i := 0; i < q; i++ {
		fmt.Println(lowerBound(nums, query[i]), upperBound(nums, query[i]))

	}

}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l < len(nums) && nums[l] == target {
		return l
	}
	return -1
}

func upperBound(nums []int, target int) int {
	l, r := -1, len(nums)-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if r > -1 && nums[r] == target {
		return r
	}
	return -1
}
