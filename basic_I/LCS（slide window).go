package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 1000010

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	nums := strings.Split(line, " ")

	var a [M]int
	for i := 0; i < n; i++ {
		a[i], _ = strconv.Atoi(nums[i])
	}
	var l, r, ans int
	window := map[int]int{}

	for r < n {

		ch1 := a[r]

		window[ch1] += 1

		for window[ch1] == 2 {
			ch2 := a[l]
			window[ch2] -= 1
			l++
		}

		ans = max(ans, r-l+1)

		r++
	}

	fmt.Println(ans)
}
