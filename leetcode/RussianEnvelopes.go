package main

import (
	"fmt"
	"sort"
)

func main() {

	var envelopes = [][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}

	ans := maxEnvelopes(envelopes)

	fmt.Println(ans)

}

// LIS 超时解法
func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	dp := make([]int, len(envelopes))
	ans := 1
	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[j][1] < envelopes[i][1] {
				dp[i] = maxT(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}

	return ans
}

func maxT(i, j int) int {

	if i > j {
		return i
	}
	return j

}

func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		// 按第一维升序 第二维降序排
		// 第二维降序的原因是防止错误计数同样宽的信封将小的替了
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	var b []int
	for i := 0; i < len(envelopes); i++ {
		p := binarySearch(b, envelopes[i][1])
		if p < len(b) {
			b[p] = envelopes[i][1]
		} else {
			b = append(b, envelopes[i][1])
		}
	}

	return len(b)
}

func binarySearch(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
