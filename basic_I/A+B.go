package main

import "fmt"

func main() {
	var A [1e6]int
	var B [1e6]int
	var C [1e6]int

	var numA string
	var numB string

	_, _ = fmt.Scanln(&numA)
	_, _ = fmt.Scanln(&numB)

	cnt1 := 0
	cnt2 := 0

	for i := len(numA) - 1; i >= 0; i-- {
		cnt1++
		A[cnt1] = int(numA[i] - '0')
	}

	for i := len(numB) - 1; i >= 0; i-- {
		cnt2++
		B[cnt2] = int(numB[i] - '0')
	}

	tot := add(A, B, C, max(cnt1, cnt2))

	for i := tot; i > 0; i-- {
		fmt.Printf("%d", C[i])
	}
}

func add(A, B, C [1e6]int, cnt int) int {
	t := 0
	for i := 1; i <= cnt; i++ {
		t += A[i] + B[i]
		C[i] = t % 10
		t /= 10
	}
	if t > 0 {
		cnt++
		C[cnt] = 1
	}
	return cnt
}
