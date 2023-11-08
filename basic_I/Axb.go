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
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}

	A = multi(A, b)
	c := len(A) - 1
	for A[c] == 0 && c > 0 {
		c--
	}
	for i := c; i >= 0; i-- {
		fmt.Printf("%d", A[i])
	}

}

func multi(A []int, b int) []int {
	// 用q存储多个进位 每次只进一个位
	q := 0
	for i := 0; i < len(A); i++ {
		q = b*A[i] + q
		A[i] = q % 10
		q /= 10
	}
	for q > 0 {
		A = append(A, q%10)
		q /= 10
	}

	return A
}

// 第二种multi写法
func multi_(A []int, b int) []int {
	var t int

	var C []int
	for i := 0; i < len(A); i++ {
		t = A[i]*b + t
		C = append(C, t%10)
		t = t / 10
	}
	//因为最后一位存的是高位 打印的时候会连续 所以一次性append就可以
	if t > 0 {
		C = append(C, t)
	}

	return C
}
