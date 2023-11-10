package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 10

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m, q int
	var matDiff [N][N]int

	_, _ = fmt.Scanln(&n, &m, &q)

	var num int
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		words := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			num, _ = strconv.Atoi(words[j])
			insert(i, j, i, j, num, &matDiff)
		}
	}
	fmt.Println(matDiff)
	var x1, y1, x2, y2, c int
	for i := 0; i <= q; i++ {
		_, _ = fmt.Scanln(&x1, &y1, &x2, &y2, &c)
		insert(x1, y1, x2, y2, c, &matDiff)
	}

	var ans [N][N]int
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			ans[i][j] = ans[i][j-1] + ans[i-1][j] - ans[i-1][j-1] + matDiff[i-1][j-1]
			b := strconv.Itoa(ans[i][j]) + " "
			_, _ = writer.WriteString(b)
		}
		_, _ = writer.WriteString("\n")
	}
	_ = writer.Flush()

}

func insert(x1, y1, x2, y2, c int, matdiff *[N][N]int) {

	matdiff[x1][y1] += c
	matdiff[x1][y2+1] -= c
	matdiff[x2+1][y1] -= c
	matdiff[x2+1][y2+1] += c

}
