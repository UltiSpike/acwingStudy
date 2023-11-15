package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 100010

var p [M]int
var l [M]int
var r [M]int

var idx = 2

// 在k的右端插入一个数
func addP(k, x int) {
	p[idx] = x
	l[idx] = k
	r[idx] = r[k]
	l[r[k]] = idx
	r[k] = idx
	idx++
}

// 移除第k个数
func remove(k int) {
	r[l[k]] = r[k]
	l[r[k]] = l[k]
}

// --- = --- = ---
//

func main() {
	// 0是最左端 链表从2开始
	l[1] = 0
	r[0] = 1
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		switch words[0] {
		// 最左端 即是在0右端插入
		case "L":
			v, _ := strconv.Atoi(words[1])
			addP(0, v)
		// 最右端 即是在1的前一个端点的右端进行插入
		case "R":
			v, _ := strconv.Atoi(words[1])
			addP(l[1], v)
		case "D":
			k, _ := strconv.Atoi(words[1])
			remove(k + 1)
		case "IL":
			k, _ := strconv.Atoi(words[1])
			v, _ := strconv.Atoi(words[2])
			addP(l[k+1], v)
		case "IR":
			k, _ := strconv.Atoi(words[1])
			v, _ := strconv.Atoi(words[2])
			addP(k+1, v)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	for i := r[0]; i != 1; i = r[i] {
		v := p[i]
		_, _ = writer.WriteString(strconv.Itoa(v) + " ")
	}
	_ = writer.Flush()

}
