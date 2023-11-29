package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const P = 1000010

var parent [P]int

func find(x int) int {
	if x != parent[x] {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")

		switch words[0] {
		case "M":
			a, _ := strconv.Atoi(words[1])
			b, _ := strconv.Atoi(words[2])
			parent[find(a)] = parent[find(b)]
		case "Q":
			a, _ := strconv.Atoi(words[1])
			b, _ := strconv.Atoi(words[2])
			if find(a) == find(b) {
				_, _ = writer.WriteString("Yes \n")
			} else {
				_, _ = writer.WriteString("No \n")
			}
		}
	}
	_ = writer.Flush()
}
