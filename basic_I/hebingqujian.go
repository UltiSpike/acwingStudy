package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	ans := n
	var area [][2]int

	var l, r int
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		l, _ = strconv.Atoi(words[0])
		r, _ = strconv.Atoi(words[1])
		area = append(area, [2]int{l, r})
	}

	sort.Slice(area, func(i, j int) bool {
		return area[i][0] < area[j][0]
	})

	last_index := area[0][0] - 1

	for _, v := range area {
		l, r = v[0], v[1]

		if l <= last_index {
			ans -= 1
		}

		if r > last_index {
			last_index = r
		}
	}

	fmt.Println(ans)

}
