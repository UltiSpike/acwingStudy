package basic_iii

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const N = 510
//
//var d [N]int
//var g [N][N]int
//var visited [N]int
//var n, m int
//
//func dijkstra() {
//	for i := 2; i <= n; i++ {
//		d[i] = 1e9
//	}
//
//	for i := 0; i < n; i++ {
//		v := -1
//		for j := 1; j <= n; j++ {
//			if visited[j] == 0 && (v == -1 || d[j] < d[v]) {
//				v = j
//			}
//		}
//		visited[v] = 1
//
//		for j := 2; j <= n; j++ {
//			d[j] = min(d[j], d[v]+g[v][j])
//		}
//
//	}
//
//}
//
//func main() {
//
//	_, _ = fmt.Scanln(&n, &m)
//
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= n; j++ {
//			g[i][j] = 1e9
//		}
//
//	}
//
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		x, _ := strconv.Atoi(words[0])
//		y, _ := strconv.Atoi(words[1])
//		w, _ := strconv.Atoi(words[2])
//		g[x][y] = min(g[x][y], w)
//	}
//
//	dijkstra()
//	if d[n] < 1e9 {
//		fmt.Println(d[n])
//	} else {
//		fmt.Println(-1)
//	}
//
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
