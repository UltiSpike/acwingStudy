package basic_iii

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var d [N]int
var m int

func bfs() int {

	if n == 1 {
		return 0
	}

	q := []int{}

	q = append(q, 1)
	var lth int
	for len(q) > 0 {
		lth = len(q)
		for i := 0; i < lth; i++ {
			t := q[0]
			for p := h[t]; p != -1; p = ne[p] {
				v := e[p]
				if d[v] == 0 {
					d[v] = d[t] + 1
					if v == n {
						return d[n]
					}
					q = append(q, v)
				}
			}
			q = q[1:]
		}

	}

	return -1
}

func main() {
	_, _ = fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i <= n; i++ {
		h[i] = -1

	}

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		add(x, y)
	}

	fmt.Println(bfs())

}
