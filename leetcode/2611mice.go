package main

import "sort"

// 等价于都是二鼠吃 然后1是选择k块收益最高的（可以用快速TopK选择划分右边区间是前K大的数，左边是前len(nums)-K小的数）
func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x // 先全部给第二只老鼠
		reward1[i] -= x
	}
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}

// 00 : 19 : 00  贪心，计算收益后排序，吃收益最高的K块
func miceAndCheese2(reward1 []int, reward2 []int, k int) int {
	gain := make([][2]int, len(reward1))
	for i := 0; i < len(reward1); i++ {
		gain[i][0] = reward1[i] - reward2[i]
		gain[i][1] = i
	}
	sort.Slice(gain, func(i, j int) bool {
		return gain[i][0] < gain[j][0]
	})
	var ans int
	for i := len(gain) - 1; i >= len(gain)-k; i-- {
		ans += reward1[gain[i][1]]
		reward2[gain[i][1]] = 0
	}
	for _, v := range reward2 {
		ans += v
	}

	return ans
}
