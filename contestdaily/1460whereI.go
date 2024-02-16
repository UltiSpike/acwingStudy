package main

import "fmt"

// 怎么求最长重复子序列

func check(mid int, s string) bool {
	table := map[string]int{}
	var u string
	for i := 0; i+mid <= len(s); i++ {
		u = s[i : i+mid]
		if table[u] > 0 {
			return true
		} else {
			table[u] = 1
		}
	}

	return false
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	var s string
	_, _ = fmt.Scanln(&s)
	l := 1
	r := n
	var mid int
	for l < r {
		mid = (l + r) >> 1
		if check(mid, s) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Println(l)
}
