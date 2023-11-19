package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const zz = 100010

var stack [zz]int
var tt int

func push(x int) {
	tt++
	stack[tt] = x
}

func pop() {
	tt--
}

func empty() string {
	if tt == 0 {
		return "YES"
	}
	return "NO"
}

func query() int {
	return stack[tt]
}

func main() {
	var M int
	_, _ = fmt.Scanln(&M)
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < M; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		operate := words[0]
		switch operate {
		case "push":
			x, _ := strconv.Atoi(words[1])
			push(x)
		case "pop":
			pop()
		case "query":
			fmt.Println(query())
		case "empty":
			fmt.Println(empty())
		}
	}

}
