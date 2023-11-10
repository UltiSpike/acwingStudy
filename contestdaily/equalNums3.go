package main

import "fmt"

const N = 1e6

// 解法1 暴力解法:过不了
//func main() {
//	var n int
//	var preSum [N]int
//	var num int
//	_, _ = fmt.Scanln(&n)
//	for i := 0; i < n; i++ {
//		_, _ = fmt.Scanf("%d", &num)
//		preSum[i+1] = num + preSum[i]
//	}
//	var sum1, sum2, sum3 int
//	var ans int
//	for i := 0; i < n; i++ {
//		for j := n - 1; j > i+1; j-- {
//			sum1 = preSum[i+1]
//			sum2 = preSum[j] - preSum[i+1]
//			sum3 = preSum[n] - preSum[j]
//			if sum1 == sum2 && sum1 == sum3 && j > i {
//				ans += 1
//			}
//		}
//	}
//
//	fmt.Println(ans)
//
//}

// 解法2
func main() {
	var n int
	var preSum [N]int
	var num int
	_, _ = fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		preSum[i+1] = num + preSum[i]
	}
	if preSum[n]%3 != 0 {
		fmt.Println("0")
		return
	}

	var avg int
	avg = preSum[n] / 3

	var sum1, sum2 int
	var ans, cnt int
	for i := 0; i < n-1; i++ {
		sum1 = preSum[i+1]
		sum2 = preSum[n] - preSum[i+1] // 控制中间区间长度>1
		if sum2 == avg {
			ans += cnt
		}
		if sum1 == avg {
			cnt++
		}
	}

	fmt.Println(ans)

}
