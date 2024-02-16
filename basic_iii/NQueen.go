package basic_iii

import (
	"bufio"
	"fmt"
	"os"
)

var col []int
var dg []int
var udg []int

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	col = make([]int, n)
	dg = make([]int, 2*n)
	udg = make([]int, 2*n)
	var path = make([][]string, n)
	for i := 0; i < n; i++ {
		path[i] = make([]string, n)
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}

	dfs(0, n, path)
}

func dfs(x, n int, path [][]string) {
	if x == n {
		writer := bufio.NewWriter(os.Stdout)
		for i := 0; i < len(path); i++ {
			for j := 0; j < len(path[0]); j++ {
				_, _ = writer.WriteString(path[i][j])
			}
			_, _ = writer.WriteString("\n")
		}
		_, _ = writer.WriteString("\n")
		_ = writer.Flush()
		return
	}

	for y := 0; y < len(path[0]); y++ {
		if col[y] == 0 && dg[x+y] == 0 && udg[n+y-x] == 0 {
			col[y] = 1
			dg[x+y] = 1
			udg[n+y-x] = 1
			path[x][y] = "Q"
			dfs(x+1, n, path)
			path[x][y] = "."
			col[y] = 0
			dg[x+y] = 0
			udg[n+y-x] = 0
		}
	}

}
