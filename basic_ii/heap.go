package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//

var heap [N]int
var size int

func up(x int) {
	if x <= 1 {
		return
	}
	if heap[x/2] > heap[x] {
		heap[x/2], heap[x] = heap[x], heap[2/x]
		up(x / 2)
	}
	return
}

func down(x int) {
	if x > size {
		return
	}
	var p = x
	if 2*x <= size && heap[2*x] < heap[p] {
		p = 2 * x
	}
	if 2*x+1 <= size && heap[2*x+1] < heap[p] {
		p = 2*x + 1
	}
	if x != p {
		heap[p], heap[x] = heap[x], heap[p]
		down(p)
	}

}

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(words[i])
		size++
		heap[size] = x
	}
	//O(n)时间建堆
	for i := n / 2; i > 0; i-- {
		down(i)
	}
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < m; i++ {
		_, _ = writer.WriteString(strconv.Itoa(heap[1]) + " ")
		heap[1] = heap[size]
		size--
		down(1)
	}

	_ = writer.Flush()
}
