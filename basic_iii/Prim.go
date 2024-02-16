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
//var g [N][N]int
//var d, st [N]int
//var n, m int
//
//func prim() {
//	var res int
//
//	for i := 1; i <= n; i++ {
//		d[i] = 1e9
//	}
//
//	for i := 0; i < n; i++ {
//		e := -1
//		for j := 1; j <= n; j++ {
//			if st[j] == 0 && (e == -1 || d[e] > d[j]) {
//				e = j
//			}
//		}
//		if i > 0 && d[e] == 1e9 {
//			fmt.Println("impossible")
//			return
//		}
//		if i > 0 {
//			res += d[e]
//		}
//
//		for j := 1; j <= n; j++ {
//			d[j] = min(d[j], g[e][j])
//		}
//		st[e] = 1
//	}
//
//	fmt.Println(res)
//	return
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
//
//func main() {
//	_, _ = fmt.Scanln(&n, &m)
//
//	for i := 1; i <= n; i++ {
//		for j := 1; j <= n; j++ {
//			if i == j {
//				continue
//			}
//			g[i][j] = 1e9
//		}
//	}
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		x, _ := strconv.Atoi(words[0])
//		y, _ := strconv.Atoi(words[1])
//		z, _ := strconv.Atoi(words[2])
//		g[x][y] = min(g[x][y], z)
//		g[y][x] = g[x][y]
//	}
//
//	prim()
//
//}
