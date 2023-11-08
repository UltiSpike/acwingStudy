package main

import "fmt"

func main() {
	var a string
	var b string

	_, _ = fmt.Scanln(&a)
	_, _ = fmt.Scanln(&b)

	var A []int
	var B []int

	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}

	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}

	flag := cmp(A, B)
	if flag == true {
		minus(A, B)
	} else {
		minus(B, A)
		A = B
		fmt.Printf("-")
	}
	i := len(A) - 1
	for A[i] == 0 && i > 0 {
		i--
	}
	for b := i; b >= 0; b-- {
		fmt.Printf("%d", A[b])
	}
}

func cmp(A, B []int) bool {
	if len(A) != len(B) {
		return len(A) > len(B)
	}
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] != B[i] {
			return A[i] > B[i]
		}
	}

	return true
}

func minus(A, B []int) {
	t := 0
	// 减法需要一直向上借位 所以i的循环长度是len(A)
	for i := 0; i < len(A); i++ {
		t = A[i] - t
		if i < len(B) {
			t -= B[i]
		}
		A[i] = (t + 10) % 10
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}

}
