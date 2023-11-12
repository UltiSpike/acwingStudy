package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	A := strings.Split(line, " ")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	B := strings.Split(line, " ")
	var l, r int
	for l < n && r < m {
		if B[r] == A[l] {
			l += 1
		}
		r++
	}

	if l == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
