package main

import (
	"fmt"
)

func main() {
	var a string
	var b int
	_, _ = fmt.Scanln(&a)
	_, _ = fmt.Scanln(&b)

	var A []int
	for i := 0; i < len(a); i++ {
		A = append(A, int(a[i]-'0'))
	}

	ANS, ans := div(A, b)
	c := 0
	for ANS[c] == 0 && c < len(A)-1 {
		c++
	}
	for _, v := range ANS[c:] {
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n%d", ans)
}

func div(A []int, b int) ([]int, int) {
	// 用q存储多个进位 每次只进一个位
	var C []int
	q := 0
	for i := 0; i < len(A); i++ {
		// 由于是上一位的余数 所以要乘10
		q = A[i] + q*10
		C = append(C, q/b)
		q = q % b
	}

	return C, q
}
