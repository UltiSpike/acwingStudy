package basic_iii

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var in [N]int

func topsort() {

	q := []int{}

	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
		}
	}

	var path = []int{}

	for len(q) > 0 {
		p := q[0]

		path = append(path, p)
		for v := h[p]; v != -1; v = ne[v] {
			t := e[v]
			in[t] -= 1
			if in[t] == 0 {
				q = append(q, t)
			}

		}

		q = q[1:]

	}

	if len(path) == n {
		writer := bufio.NewWriter(os.Stdout)
		for i := 0; i < n; i++ {
			_, _ = writer.WriteString(strconv.Itoa(path[i]) + " ")
		}
		_ = writer.Flush()
		return
	} else {
		fmt.Println(-1)
	}

	return
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

	topsort()

}
