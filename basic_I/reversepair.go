package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d \n", &n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &nums[i])
	}
	ans := reversepair(nums, 0, n-1)

	fmt.Println(ans)
}

func reversepair(nums []int, l, r int) int {
	if l >= r {
		return 0
	}
	var ans int
	mid := (l + r) >> 1
	left := reversepair(nums, l, mid)
	right := reversepair(nums, mid+1, r)
	var temp []int
	i, j := l, mid+1
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else if nums[i] > nums[j] {
			temp = append(temp, nums[j])
			ans += mid - i + 1
			j++
		}
	}
	for i <= mid {
		temp = append(temp, nums[i])
		i++
	}

	for j <= r {
		temp = append(temp, nums[j])
		j++
	}

	for p := l; p <= r; p++ {
		nums[p] = temp[p-l]
	}

	return ans + left + right
}
