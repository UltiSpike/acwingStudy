package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1000010

var p [N]int
var cnt [N]int

func find(x int) int {
	if x != p[x] {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)

	for i := 1; i <= n; i++ {
		p[i] = i
		cnt[i] = 1
	}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		switch words[0] {
		case "C":
			a, _ := strconv.Atoi(words[1])
			b, _ := strconv.Atoi(words[2])
			if find(a) == find(b) {
				continue
			}
			cnt[find(b)] += cnt[find(a)]
			p[find(a)] = p[find(b)]

		case "Q1":
			a, _ := strconv.Atoi(words[1])
			b, _ := strconv.Atoi(words[2])
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		case "Q2":
			a, _ := strconv.Atoi(words[1])
			fmt.Println(cnt[find(a)])
		}

	}

}
