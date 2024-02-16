package basic_iii

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

// 邻接表形式存图

// 地址对应的节点
var e [N]int

// 下一个节点地址
var ne [N]int

// 节点的邻接表
var h [N]int
var idx int
var n int
var ans int
var visited [N]int

func add(x, y int) {
	idx++
	e[idx] = x
	ne[idx] = h[y]
	h[y] = idx

}

func main() {

	_, _ = fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		h[i] = -1
	}

	reader := bufio.NewReader(os.Stdin)
	ans = n
	for i := 0; i < n-1; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")

		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])

		//fmt.Println(x, y)
		add(x, y)
		add(y, x)

	}

	backtrace(1)
	fmt.Println(ans)
}

func backtrace(root int) int {

	visited[root] = 1
	var son, temp int
	var res int
	var other int

	for p := h[root]; p != -1; p = ne[p] {
		if visited[e[p]] == 0 {
			temp = backtrace(e[p])
			son += temp
			if temp > res {
				res = temp
			}
		}

	}

	other = n - son - 1
	if other > res {
		res = other
	}

	if res < ans {
		ans = res
	}
	fmt.Println(ans, res)
	return son + 1
}
