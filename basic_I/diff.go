package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)
	var diff [N]int
	var nums [N]int

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")
	for i := 1; i <= n; i++ {

		nums[i], _ = strconv.Atoi(words[i-1])
		diff[i] = nums[i] - nums[i-1]
	}

	var l, r, c int

	for i := 0; i < m; i++ {
		inline, _ := reader.ReadString('\n')
		inline = strings.TrimSpace(inline)
		inWords := strings.Split(inline, " ")
		l, _ = strconv.Atoi(inWords[0])
		r, _ = strconv.Atoi(inWords[1])
		c, _ = strconv.Atoi(inWords[2])

		diff[l] += c
		diff[r+1] -= c
	}

	writer := bufio.NewWriter(os.Stdout)
	var ans [N]int
	for i := 1; i <= n; i++ {

		ans[i] = ans[i-1] + diff[i]
		b := strconv.Itoa(ans[i]) + " "
		_, _ = writer.WriteString(b)
	}
	_ = writer.Flush()
}

//
//1 7 2 4
//3 6 2 8
//2 1 2 3
