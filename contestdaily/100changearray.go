package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 10分钟偷看答案

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	nums := make([]int, n+1)
	diff := make([]int, n+1)

	for i := 1; i <= n; i++ {

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		num, _ := strconv.Atoi(words[0])
		nums[i] = num
		diff[i-1] = nums[i] - nums[i-1]

	}

	var count, ans int

	count, ans = counter(diff[1:n])

	fmt.Println(count)
	fmt.Println(ans)

}

// 原问题转化为使区间内下标都为0
// 可以让一个数+1 一个数-1
// 先排序
// 相当于把(0, n-1)内化成 00000x0000的形式
// 然后还有x种

func counter(nums []int) (count, ans int) {

	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	for l < r && (nums[l] < 0 || nums[r] > 0) {
		if nums[l] == 0 {
			l++
		}
		if nums[r] == 0 {
			r--
		}
		nums[l] += 1
		nums[r] -= 1
		count++
	}
	for l <= r {
		ans += nums[l]
		l++
	}
	if ans < 0 {
		ans = -ans
	}

	return count, ans
}
