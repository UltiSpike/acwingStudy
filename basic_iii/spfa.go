package basic_iii

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const N = 100010
//
//var e, ne, h, w, d, st [N]int
//var idx int
//var n, m int
//
//func add(x, y, z int) {
//	idx++
//	e[idx] = y
//	ne[idx] = h[x]
//	h[x] = idx
//	w[idx] = z
//
//}
//
//func spfa() {
//	for i := 2; i <= n; i++ {
//		d[i] = 1e9
//	}
//
//	var q []int
//	q = append(q, 1)
//	st[1] = 1
//	for len(q) > 0 {
//
//		t := q[0]
//		q = q[1:]
//		st[t] = 0
//		for p := h[t]; p != -1; p = ne[p] {
//			E := e[p]
//
//			if d[E] > d[t]+w[p] {
//				d[E] = d[t] + w[p]
//				if st[E] == 0 {
//					q = append(q, E)
//					st[E] = 1
//				}
//
//			}
//
//		}
//
//	}
//
//}
//
//func min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//
//}
//
//func main() {
//	_, _ = fmt.Scanln(&n, &m)
//
//	for i := 1; i <= n; i++ {
//		h[i] = -1
//	}
//
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		a, _ := strconv.Atoi(words[0])
//		b, _ := strconv.Atoi(words[1])
//		z, _ := strconv.Atoi(words[2])
//		add(a, b, z)
//	}
//
//	spfa()
//
//	if d[n] >= 1e9/2 {
//		fmt.Println("impossible")
//	} else {
//		fmt.Println(d[n])
//	}
//
//}
