package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	var num int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		fmt.Printf("%d", conuntbye(num))

	}
}

func conuntbye(num int) int {
	ans := 0
	for num > 0 {
		ans += num & 1 // num & 1 结果是 整型的1或 0
		num /= 2
	}

	return ans

}
