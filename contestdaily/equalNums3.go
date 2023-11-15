package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	var num int
	var a [N]int
	for i := 0; i < n; i++ {
		num, _ = strconv.Atoi(words[i])
		a[i+1] = a[i] + num
	}
	// 小心数组求和/3 不整除的情况
	if a[n]%3 != 0 {
		fmt.Println("0")
		return
	}

	avg := a[n] / 3
	var preSum, suffixSum int
	var cnt int
	var ans int
	// 左边界取值[0, n-2] 右边界取值[1,n-1]
	for i := 0; i < n-1; i++ {
		preSum = a[i+1]           // a[i] 前i项的前缀和和  a[i+1] 代表[0, i]的和
		suffixSum = a[n] - a[i+1] // a[r]- a[l-1],区间(l, r)的和 a[n]- a[i+1] 代表 (i, n-1] 的前缀和
		// 对于当前的r 合法的 L 处在 [0, i)区间里
		// 所以不计当前i位置的cnt
		if suffixSum == avg {
			ans += cnt
		}
		if preSum == avg {
			cnt++
		}
	}

	fmt.Println(ans)
}
