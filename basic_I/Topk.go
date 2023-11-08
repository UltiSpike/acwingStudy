package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, k int
	_, _ = fmt.Scanf("%d %d \n", &n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}

	ans := findKth(nums, k)
	fmt.Println(ans)
}

func findKth(nums []int, k int) int {

	pivot_order := index(nums) + 1

	if pivot_order == k {
		return nums[pivot_order-1]
	}
	if pivot_order < k {

		return findKth(nums[pivot_order:], k-pivot_order)
	} else {

		return findKth(nums[:pivot_order-1], k)
	}
}

func index(nums []int) int {
	pivot := rand.Intn(len(nums))
	nums[pivot], nums[0] = nums[0], nums[pivot]
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[i], nums[j+1] = nums[j+1], nums[i]
			j++
		}
	}
	nums[j], nums[0] = nums[0], nums[j]
	return j
}

//
