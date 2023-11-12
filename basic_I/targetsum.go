package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Z = 1000010

func main() {
	var n, m, x int
	_, _ = fmt.Scanln(&n, &m, &x)

	reader := bufio.NewReader(os.Stdin)
	var A [Z]int
	var B [Z]int
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	for i := 0; i < n; i++ {
		A[i], _ = strconv.Atoi(words[i])
	}

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words = strings.Split(line, " ")
	for i := 0; i < m; i++ {
		B[i], _ = strconv.Atoi(words[i])
	}

	l, r := 0, m-1

	for l < n && r >= 0 {

		if A[l]+B[r] == x {
			fmt.Println(l, r)
			return
		}
		if A[l]+B[r] < x {
			l++
		} else if A[l]+B[r] > x {
			r--
		}

	}

}
