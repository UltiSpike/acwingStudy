package main

import (
	"fmt"
)

func main() {
	var num float64
	_, _ = fmt.Scanf("%f", &num)
	fmt.Printf("%.6f", cuberRoot(num))
}

func cuberRoot(num float64) float64 {
	var l, r float64
	l = -1000
	r = 1000
	for r-l > 1e-7 {
		mid := (l + r) / 2
		if mid*mid*mid < num {
			l = mid
		} else {
			r = mid
		}
	}

	return l
}

// 浮点数二分
// 对于浮点数二分而言，“求分界点” 也就是找一个能够满足精度要求的解就可以了
// 浮点数二分先根据数据的范围来确定二分答案的区间的范围[ l , r ] [l, r][l,r]，
// 浮点数二分就是对求解答案来二分，所以区间长度r - l就约束了答案的精度。
