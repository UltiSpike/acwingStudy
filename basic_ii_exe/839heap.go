package basic_ii_exe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

var heap [N]int
var size = 0
var len = 0
var ph [N]int
var hp [N]int

func heap_swap(x, y int) {

	heap[x], heap[y] = heap[y], heap[x]
	hp[x], hp[y] = hp[y], hp[x]
	ph[hp[x]], ph[hp[y]] = ph[hp[y]], ph[hp[x]]
}

func up(k int) {
	if k/2 > 0 && heap[k/2] > heap[k] {
		heap_swap(k/2, k)
		up(k / 2)
	}
	return
}

func down(x int) {
	var t = x
	if x*2 <= size && heap[2*x] < heap[t] {
		t = 2 * x
	}
	if x*2+1 <= size && heap[2*x+1] < heap[t] {
		t = 2*x + 1
	}
	if t != x {
		heap_swap(t, x)
		down(t)
	}
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		switch words[0] {
		case "I":
			x, _ := strconv.Atoi(words[1])
			size++
			len++
			heap[size] = x
			ph[len] = size
			hp[size] = len
			up(size)
		case "PM":
			_, _ = writer.WriteString(strconv.Itoa(heap[1]) + "\n")
		case "DM":
			heap_swap(1, size)
			size--
			down(1)
		case "C":
			k, _ := strconv.Atoi(words[1])
			x, _ := strconv.Atoi(words[2])
			k = ph[k]
			heap[k] = x
			up(k)
			down(k)
		case "D":
			k, _ := strconv.Atoi(words[1])
			k = ph[k]
			heap_swap(k, size)
			size--
			down(k)
			up(k)

		}

	}
	_ = writer.Flush()
}
