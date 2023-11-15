package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Nn = 1010

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, m, q int
	var b [Nn][Nn]int

	_, _ = fmt.Scanln(&n, &m, &q)

	var num int
	for i := 1; i <= n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		for j := 1; j <= m; j++ {
			num, _ = strconv.Atoi(words[j-1])
			insert(i, j, i, j, num, &b)
		}

	}

	var x1, y1, x2, y2, c int
	for i := 0; i < q; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		x1, _ = strconv.Atoi(words[0])
		y1, _ = strconv.Atoi(words[1])
		x2, _ = strconv.Atoi(words[2])
		y2, _ = strconv.Atoi(words[3])
		c, _ = strconv.Atoi(words[4])

		insert(x1, y1, x2, y2, c, &b)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]
			c := strconv.Itoa(b[i][j]) + " "
			_, _ = writer.WriteString(c)
		}
		_, _ = writer.WriteString("\n")
	}
	_ = writer.Flush()

}

func insert(x1, y1, x2, y2, c int, matdiff *[Nn][Nn]int) {

	matdiff[x1][y1] += c
	matdiff[x1][y2+1] -= c
	matdiff[x2+1][y1] -= c
	matdiff[x2+1][y2+1] += c

}

// preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + nums[i][j]
// nums[i][j] = preSum[i][j] - (preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] )
