package main

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
//var heap, hp, ph [N]int
//var size, len int
//
//// hp 堆中节点代表第几次插入
//// ph 第几次插入指向的堆中节点
//
//func swap(x, y int) {
//	heap[x], heap[y] = heap[y], heap[x]
//	hp[x], hp[y] = hp[y], hp[x]
//	// 反正还是他们俩 能交换出来
//	ph[hp[x]], ph[hp[y]] = ph[hp[y]], ph[hp[x]]
//}
//
//// 将x节点下沉
//func down(x int) {
//	if x > size {
//		return
//	}
//	var t = x
//	if 2*x <= size && heap[2*x] < heap[t] {
//		t = 2 * x
//	}
//
//	if 2*x+1 <= size && heap[2*x+1] < heap[t] {
//		t = 2*x + 1
//	}
//
//	if t != x {
//		swap(x, t)
//
//		down(t)
//	}
//}
//
//func up(x int) {
//	if x <= 1 {
//		return
//	}
//	if heap[x/2] > heap[x] {
//		swap(x, x/2)
//		up(x / 2)
//	}
//
//}
//
//func main() {
//	var n int
//	_, _ = fmt.Scanln(&n)
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	for i := 0; i < n; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		switch words[0] {
//		case "I":
//			x, _ := strconv.Atoi(words[1])
//			size++
//			len++
//			heap[size] = x
//			ph[len] = size
//			hp[size] = len
//			up(size)
//		case "PM":
//			_, _ = writer.WriteString(strconv.Itoa(heap[1]) + "\n")
//		case "DM":
//			swap(1, size)
//			size--
//			down(1)
//		case "C":
//			k, _ := strconv.Atoi(words[1])
//			x, _ := strconv.Atoi(words[2])
//			k = ph[k]
//			heap[k] = x
//			up(k)
//			down(k)
//		case "D":
//			k, _ := strconv.Atoi(words[1])
//			k = ph[k]
//			swap(k, size)
//			size--
//			down(k)
//			up(k)
//
//		}
//
//	}
//	_ = writer.Flush()
//}
