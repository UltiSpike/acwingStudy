package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const N = 300010

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	// 一个数组存用过的下标
	var a []int
	// 一个数组存操作
	var ad [][2]int
	// 一个数组存离散后操作完的数组和计算前缀和
	var b [N]int
	// 一个数组存查询
	var query [][2]int
	var x, c int
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		x, _ := strconv.Atoi(words[0])
		c, _ := strconv.Atoi(words[1])
		a = append(a, x)
		inline := [2]int{x, c}
		ad = append(ad, inline)
	}

	var l, r int
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		l, _ := strconv.Atoi(words[0])
		r, _ := strconv.Atoi(words[1])
		inline := [2]int{l, r}
		query = append(query, inline)
		a = append(a, l)
		a = append(a, r)
	}
	sort.Ints(a)

	a = removeDuplicates(a)
	for i := 0; i < n; i++ {
		x, c = ad[i][0], ad[i][1]
		b[find(x, a)] += c
	}
	for i := 1; i <= len(a); i++ {
		b[i] += b[i-1]
	}

	for i := 0; i < m; i++ {
		l, r = query[i][0], query[i][1]
		fmt.Println(b[find(r, a)] - b[find(l, a)-1])

	}

}

func find(v int, nums []int) int {
	l := 0
	r := len(nums)
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nums[mid] < v {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l + 1
}

func removeDuplicates(nums []int) []int {
	curentIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[curentIndex] {
			curentIndex++
			nums[curentIndex] = nums[i]
		}
	}

	return nums[:curentIndex+1]
}
