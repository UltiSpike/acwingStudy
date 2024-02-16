package basic_ii_exe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 31 * 100010

// Node M个节点数，一个节点存字符串的一位字符
// 数字的大小是2的31次方，代表一个数字需要31个节点来存，一共有N个数字，所以一共需要31*N个节点
var Node [M][2]int
var idx int

func Insert(x int) {

	var p int
	var u int

	for i := 30; i >= 0; i-- {
		u = x >> i & 1
		if Node[p][u] == 0 {
			idx++
			Node[p][u] = idx
		}
		p = Node[p][u]
	}

}

func query(x int) int {
	var res int

	var p int
	var u int
	var t int
	for i := 30; i >= 0; i-- {
		u = x >> i & 1
		if u == 1 {
			t = 0
		} else {
			t = 1
		}
		if Node[p][t] != 0 {
			res = res*2 + 1
			p = Node[p][t]
		} else {
			res = res * 2
			p = Node[p][u]
		}

	}

	return res
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(words[i])

		Insert(num)
	}

	var ans int
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(words[i])

		ans = max(ans, query(num))
	}

	fmt.Println(ans)

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
