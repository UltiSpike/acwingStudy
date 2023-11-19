package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const que = 100000

var queue [que]int
var start = 1
var end int

func pushQ(x int) {
	end++
	queue[end] = x
}

func isQEmpty() string {
	if end < start {
		return "YES"
	}
	return "NO"
}

func popQ() {
	start++
}

func queryQ() int {
	return queue[start]
}

func main() {
	var M int
	_, _ = fmt.Scanln(&M)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		switch words[0] {
		case "push":
			x, _ := strconv.Atoi(words[1])
			pushQ(x)
		case "query":
			x := queryQ()
			fmt.Println(x)
		case "empty":
			fmt.Println(isQEmpty())
		case "pop":
			popQ()

		}
	}

}
