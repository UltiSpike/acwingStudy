package basic_iii

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//const N = 100010
//
//type edge struct {
//	a int
//	b int
//	v int
//}
//
//var edges []edge
//var p [N]int
//
//func find(x int) int {
//	if p[x] != x {
//		p[x] = find(p[x])
//	}
//
//	return p[x]
//}
//
//var n, m int
//
//func kruskal() {
//	sort.Slice(edges, func(i, j int) bool {
//		return edges[i].v < edges[j].v
//	})
//
//	var res int
//	var cnt int
//	for i := 0; i < m; i++ {
//		a := find(edges[i].a)
//		b := find(edges[i].b)
//		v := edges[i].v
//
//		if a != b {
//			p[a] = b
//			res += v
//			cnt++
//		}
//	}
//
//	if cnt < n-1 {
//		fmt.Println("impossible")
//	} else {
//		fmt.Println(res)
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
//
//func main() {
//	_, _ = fmt.Scanln(&n, &m)
//
//	reader := bufio.NewReader(os.Stdin)
//
//	for i := 1; i <= n; i++ {
//		p[i] = i
//	}
//
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		x, _ := strconv.Atoi(words[0])
//		y, _ := strconv.Atoi(words[1])
//		z, _ := strconv.Atoi(words[2])
//		edges = append(edges, edge{x, y, z})
//
//	}
//
//	kruskal()
//
//}
