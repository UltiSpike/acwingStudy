package basic_iii

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	var martix = make([][]int, n)
	var visited = make([][]int, n)
	var num int
	for i := 0; i < n; i++ {
		visited[i] = make([]int, m)
		martix[i] = make([]int, m)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			num, _ = strconv.Atoi(words[j])
			martix[i][j] = num
		}
	}

	var q = [][2]int{}
	q = append(q, [2]int{0, 0})
	var p [2]int
	var x, y int
	var step, length int
	var dx = []int{1, 0, -1, 0}
	var dy = []int{0, 1, 0, -1}
	for len(q) > 0 {

		length = len(q)
		for i := 0; i < length; i++ {
			p = q[0]
			if p == [2]int{n - 1, m - 1} {
				fmt.Println(step)
				return
			}

			for j := 0; j < 4; j++ {
				x, y = p[0]+dx[j], p[1]+dy[j]
				if x >= 0 && y >= 0 && x < n && y < m && martix[x][y] == 0 && visited[x][y] == 0 {
					//fmt.Println(step, x, y)
					q = append(q, [2]int{x, y})
					visited[p[0]][p[1]] = 1
				}
			}
			q = q[1:]
		}
		step += 1
	}
}
