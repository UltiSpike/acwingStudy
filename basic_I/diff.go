package main

import "fmt"

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	var diff [N]int
	var nums [N]int
	var num int

	for i := 1; i <= n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		nums[i] = num
		diff[i] = nums[i] - nums[i-1]
	}

	var l, r, c int

	for i := 0; i < m; i++ {

		_, _ = fmt.Scanln(&l, &r, &c)
		diff[l] += c
		diff[r+1] -= c
	}

	var ans [N]int
	for i := 1; i <= n; i++ {

		ans[i] = ans[i-1] + diff[i]

		fmt.Printf("%d ", ans[i])
	}

}

//
//1 7 2 4
//3 6 2 8
//2 1 2 3
