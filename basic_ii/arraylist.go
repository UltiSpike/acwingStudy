package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1000010

var e [N]int
var ne [N]int
var index = 0
var head = -1

func addPointsHead(x int) {
	e[index] = x
	ne[index] = head
	head = index
	index++
}

func addPoints(k, x int) {
	e[index] = x
	ne[index] = ne[k]
	ne[k] = index
	index++
}

func deletePoints(k int) {
	ne[k] = ne[ne[k]]
}

func main() {
	var m int
	_, _ = fmt.Scanln(&m)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		switch words[0] {
		case "H":
			v, _ := strconv.Atoi(words[1])
			addPointsHead(v)
		case "D":
			k, _ := strconv.Atoi(words[1])
			if k == 0 {
				head = ne[head]
			} else {
				deletePoints(k - 1)
			}

		case "I":
			k, _ := strconv.Atoi(words[1])
			v, _ := strconv.Atoi(words[2])
			addPoints(k-1, v)
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	// 末尾是空节点 别写错了
	for t := head; t != -1; t = ne[t] {
		v := e[t]
		_, _ = writer.WriteString(strconv.Itoa(v) + " ")
	}
	_ = writer.Flush()
}
