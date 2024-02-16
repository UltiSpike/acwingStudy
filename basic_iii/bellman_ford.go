package basic_iii

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//type line struct {
//	a int
//	b int
//	w int
//}
//
//var lines []line
//
////var d []int
////var n, m, k int
//
//func bellman_ford() {
//	for i := 2; i <= n; i++ {
//		d[i] = 1e9
//	}
//	var backup = make([]int, n+1)
//	for i := 0; i < k; i++ {
//		// 必须用copy，不然是同一个切片
//		copy(backup, d)
//		for j := 0; j < m; j++ {
//			a := lines[j].a
//			b := lines[j].b
//			w := lines[j].w
//			if d[b] > backup[a]+w {
//				d[b] = backup[a] + w
//			}
//
//		}
//
//	}
//
//}
//
//func main() {
//	_, _ = fmt.Scanln(&n, &m, &k)
//	d = make([]int, n+1)
//	lines = make([]line, m+1)
//	reader := bufio.NewReader(os.Stdin)
//	for i := 0; i < m; i++ {
//		words, _ := reader.ReadString('\n')
//		words = strings.TrimSpace(words)
//		path := strings.Split(words, " ")
//		a, _ := strconv.Atoi(path[0])
//		b, _ := strconv.Atoi(path[1])
//		w, _ := strconv.Atoi(path[2])
//		lines[i] = line{a, b, w}
//
//	}
//
//	bellman_ford()
//
//	if d[n] > 1e9/2 {
//		fmt.Println("impossible")
//	} else {
//		fmt.Println(d[n])
//	}
//
//}
