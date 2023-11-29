package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	p, _ := reader.ReadString('\n')
	p = " " + strings.TrimSpace(p)

	m1, _ := reader.ReadString('\n')
	m1 = strings.TrimSpace(m1)
	m, _ := strconv.Atoi(m1)

	s, _ := reader.ReadString('\n')
	s = " " + strings.TrimSpace(s)

	next := make([]int, n+1)

	for i, j := 2, 0; i <= n; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}

	for i, j := 1, 0; i <= m; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == n {
			j = next[j]
			fmt.Printf("%d ", i-n)
		}
	}

}
