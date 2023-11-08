package main

import "fmt"

func main() {
	var n int
	var m int
	var t int
	_, _ = fmt.Scanf("%d%d\n", &n, &m)
	var pre []int
	pre = append(pre, 0)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &t)
		pre = append(pre, pre[len(pre)-1]+t)
	}
	var query [][2]int
	var inline [2]int
	var l, r int
	for i := 0; i < m; i++ {
		_, _ = fmt.Scanf("\n%d %d", &l, &r)
		inline = [2]int{l, r}
		query = append(query, inline)
	}

	// sum[l, r]
	// pre[i+1] - pre[i-1] = sum[i,i+1]
	for _, v := range query {
		l, r = v[0], v[1]
		fmt.Println(pre[r] - pre[l-1])
	}
}
