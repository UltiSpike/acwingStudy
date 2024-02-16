package basic_iii

//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const N = 1e5 + 10
//const M = 2 * N
//
//var e, ne [M]int
//var color, st, h [N]int
//var idx int
//
//func add(x, y int) {
//	idx++
//	e[idx] = y
//	ne[idx] = h[x]
//	h[x] = idx
//}
//
//var n, m int
//
//func draw(x, col int) bool {
//	color[x] = col
//	for i := h[x]; i != -1; i = ne[i] {
//		p := e[i]
//		if color[p] == col {
//			return false
//		} else if color[p] == 0 {
//			if !draw(p, 3-col) {
//				return false
//			}
//		}
//	}
//
//	return true
//}
//
//func main() {
//	_, _ = fmt.Scanln(&n, &m)
//	reader := bufio.NewReader(os.Stdin)
//	for i := 1; i <= n; i++ {
//		h[i] = -1
//	}
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		x, _ := strconv.Atoi(words[0])
//		y, _ := strconv.Atoi(words[1])
//		add(x, y)
//		add(y, x)
//	}
//
//	for i := 1; i <= n; i++ {
//		if color[i] == 0 {
//			if !draw(i, 1) {
//				fmt.Println("No")
//				return
//			}
//		}
//	}
//	fmt.Println("Yes")
//}
