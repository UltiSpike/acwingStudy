package basic_ii_exe

//
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
//var p [N]int
//var cnt [N]int
//
//func find(x int) int {
//	if p[x] != x {
//		p[x] = find(p[x])
//	}
//	return p[x]
//}
//
//func main() {
//	var n, m int
//	_, _ = fmt.Scanln(&n, &m)
//	for i := 0; i <= n; i++ {
//		p[i] = i
//		cnt[i] = 1
//	}
//
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	for i := 0; i < m; i++ {
//
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		a, _ := strconv.Atoi(words[1])
//
//		switch words[0] {
//		case "C":
//			b, _ := strconv.Atoi(words[2])
//			if find(a) == find(b) {
//				continue
//			}
//			cnt[find(b)] += cnt[find(a)]
//			p[find(a)] = find(b)
//		case "Q1":
//			b, _ := strconv.Atoi(words[2])
//			if find(a) == find(b) {
//				_, _ = writer.WriteString("Yes\n")
//			} else {
//				_, _ = writer.WriteString("No\n")
//			}
//		case "Q2":
//
//			_, _ = writer.WriteString(strconv.Itoa(cnt[find(a)]) + "\n")
//
//		}
//	}
//
//
//
//	_ = writer.Flush()
//
//}
