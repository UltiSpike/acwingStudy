package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 先想暴力

// 3 1 5
// 3 1 2 5
// 1 2

const so = 10000

var sStack [so]int
var ss int

func queryS() int {
	return sStack[ss]
}

func pushS(x int) {
	ss++
	sStack[ss] = x
}

func popS() {
	ss--
}

func main() {

	var n int
	_, _ = fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(words[i])
		z, _ := strconv.Atoi(words[queryS()])

		for ss > 0 && num <= z {
			popS()
			z, _ = strconv.Atoi(words[queryS()])
		}

		if ss > 0 {
			_, _ = writer.WriteString(strconv.Itoa(z) + " ")
		} else {
			_, _ = writer.WriteString("-1 ")
		}
		pushS(i)

	}
	_ = writer.Flush()
}
