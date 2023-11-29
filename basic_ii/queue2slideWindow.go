package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, k int
	_, _ = fmt.Scanln(&n, &k)
	maxN := make([]int, n)
	minN := make([]int, n)
	q := make([][2]int, 0)
	q2 := make([][2]int, 0)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")

	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(words[i])
		for len(q) > 0 && i-k+1 > q[0][0] {
			q = q[1:]

		}
		for len(q2) > 0 && i-k+1 > q2[0][0] {
			q2 = q2[1:]

		}

		for len(q) > 0 && num < q[len(q)-1][1] {
			q = q[:len(q)-1]
		}
		for len(q2) > 0 && num > q2[len(q2)-1][1] {
			q2 = q2[:len(q2)-1]
		}
		q = append(q, [2]int{i, num})
		q2 = append(q2, [2]int{i, num})
		minN[i] = q[0][1]
		maxN[i] = q2[0][1]
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := k - 1; i < n; i++ {
		_, _ = writer.WriteString(strconv.Itoa(minN[i]) + " ")
	}
	_, _ = writer.WriteString("\n")
	for i := k - 1; i < n; i++ {
		_, _ = writer.WriteString(strconv.Itoa(maxN[i]) + " ")
	}

	_ = writer.Flush()
}
