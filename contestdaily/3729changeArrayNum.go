package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 区间合并+差分
// 合并所有涉及到的操作区间
// 最后将涉及到的操作区间进行+1操作

func main() {
	var T int
	_, _ = fmt.Scanln(&T)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < T; i++ {
		firstLine, _ := reader.ReadString('\n')
		secondLine, _ := reader.ReadString('\n')
		firstLine = strings.TrimSpace(firstLine)
		secondLine = strings.TrimSpace(secondLine)
		first := strings.Split(firstLine, " ")
		second := strings.Split(secondLine, " ")

		b := result(first, second)
		for _, v := range b {
			_, _ = writer.WriteString(strconv.Itoa(v) + " ")
		}
		_, _ = writer.WriteString("\n")

	}

	_ = writer.Flush()

}

// 从后面往前扫 维护状态写法
// 本质也是区间合并
func trickResult(fi, se []string) []int {
	n, _ := strconv.Atoi(fi[0])

	ans := make([]int, n)
	var flag int
	for i := n - 1; i >= 0; i-- {
		t, _ := strconv.Atoi(se[i])
		if t > flag {
			flag = t
		}
		if flag > 0 {
			ans[i] = 1
			flag--
		}
	}

	return ans
}

// 差分写法
// 用差分来更新某个区间内的操作状态
// 答案和某个位置的操作状态有关
// 操作过程以区间进行
func diffResult(fi, se []string) []int {
	n, _ := strconv.Atoi(fi[0])

	ans := make([]int, n)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		r := i
		l, _ := strconv.Atoi(se[i])
		if r-l+1 < 0 {
			l = 0
		} else {
			l = r - l + 1
		}
		b[r+1] -= 1
		b[l] += 1
	}
	var cnt int
	for i := 0; i < n; i++ {
		cnt += b[i]
		if cnt > 0 {
			ans[i] = 1
		}
	}

	return ans
}

// 老麻烦的区间合并写法
// 区间合并（NLogN）
// 区间合并去除重复操作
func result(fi, se []string) []int {
	n, _ := strconv.Atoi(fi[0])

	ans := make([]int, n+1)
	var oriAdd [][2]int
	for i := 0; i < n; i++ {
		r := i
		l, _ := strconv.Atoi(se[i])
		if r-l+1 < 0 {
			l = 0
		} else {
			l = r - l + 1
		}
		if r >= l {
			oriAdd = append(oriAdd, [2]int{l, r})
		}

	}
	if len(oriAdd) <= 0 {
		return ans[:n]
	}
	// 	fmt.Println(oriAdd)
	sort.Slice(oriAdd, func(i, j int) bool {
		return oriAdd[i][0] < oriAdd[j][0]
	})

	var add [][2]int
	// 第一个区间
	var lastL = oriAdd[0][0]
	var lastR = oriAdd[0][1]
	for i := 1; i < len(oriAdd); i++ {

		if oriAdd[i][0] > lastR {
			add = append(add, [2]int{lastL, lastR})
			lastL = oriAdd[i][0]
			lastR = oriAdd[i][1]
		} else {
			if oriAdd[i][1] > lastR {
				lastR = oriAdd[i][1]
			}

		}
	}

	add = append(add, [2]int{lastL, lastR})

	for _, v := range add {
		for i := v[0]; i < v[1]+1; i++ {
			ans[i] = 1
		}
	}
	return ans[:n]
}
