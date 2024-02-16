package basic_iii

//
//import (
//	"bufio"
//	"container/heap"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//const N = 150010
//
//var e, ne, h, w, st [N]int
//var d []int
//var idx int
//
//func add2(x, y, z int) {
//	idx++
//	e2[idx] = y
//	ne[idx] = h[x]
//	h[x] = idx
//	w[idx] = z // x连到y的距离, 一个idx存的是一个距离
//}
//
//type Pair struct {
//	Distance int
//	Elements int
//}
//
//type PairHeap []Pair
//
//func (p PairHeap) Less(i, j int) bool {
//	return p[i].Distance < p[j].Distance
//}
//
//func (p PairHeap) Len() int {
//	return len(p)
//}
//
//func (p PairHeap) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//
//func (p *PairHeap) Pop() interface{} {
//	n := len(*p)
//	x := (*p)[n-1]
//	*p = (*p)[:n-1]
//
//	return x
//}
//
//func (p *PairHeap) Push(x interface{}) {
//	*p = append(*p, x.(Pair))
//
//}
//
//func dijkstra() int {
//
//	for i := 2; i <= n; i++ {
//		d[i] = 1e9
//	}
//
//	var myheap = &PairHeap{}
//	heap.Init(myheap)
//	heap.Push(myheap, Pair{0, 1})
//
//	for 0 < len(*myheap) {
//		t := heap.Pop(myheap).(Pair)
//		E := t.Elements
//		dis := t.Distance
//		if st[E] != 0 {
//			continue
//		}
//		st[E] = 1
//
//		for p := h[E]; p != -1; p = ne[p] {
//			vi := e2[p]
//			if d[vi] > dis+w[p] {
//				d[vi] = dis + w[p]
//				heap.Push(myheap, Pair{d[vi], vi})
//			}
//		}
//
//	}
//
//	return d[n]
//}
//
//var n, m int
//
//func main() {
//
//	_, _ = fmt.Scanln(&n, &m)
//	d = make([]int, n+1)
//
//	reader := bufio.NewReader(os.Stdin)
//
//	for i := 1; i <= n; i++ {
//		h[i] = -1
//	}
//	for i := 0; i < m; i++ {
//		line, _ := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		words := strings.Split(line, " ")
//		x, _ := strconv.Atoi(words[0])
//		y, _ := strconv.Atoi(words[1])
//		z, _ := strconv.Atoi(words[2])
//		add(x, y, z)
//	}
//
//	if dijkstra() == 1e9 {
//		fmt.Println(-1)
//	} else {
//		fmt.Println(d[n])
//	}
//}
