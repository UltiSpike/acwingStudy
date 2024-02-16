package basic_iii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var visitede map[string]int

func main() {
	var start string

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	words := strings.Split(line, " ")

	start = strings.Join(words, "")

	visitede = map[string]int{}

	fmt.Println(bfs3(start))

}

func bfs3(start string) int {

	var target = "12345678x"

	if start == target {
		return 0
	}
	var q = [][]byte{
		[]byte(start),
	}

	visitede[start] = 1

	var dx = []int{1, -1, 0, 0}
	var dy = []int{0, 0, 1, -1}

	for len(q) > 0 {
		p := q[0]
		x, y := find(string(p))

		k := x*3 + y
		for i := 0; i < 4; i++ {
			a := x + dx[i]
			b := y + dy[i]
			if a < 0 || a > 2 || b < 0 || b > 2 {
				continue
			}

			z := a*3 + b
			temp := make([]byte, len(p))
			copy(temp, p)
			temp[z], temp[k] = temp[k], temp[z]
			state := string(temp)

			if visitede[state] == 0 {
				q = append(q, temp)

				visitede[state] = visitede[string(p)] + 1
				if state == target {
					return visitede[state] - 1
				}
			}
		}
		q = q[1:]

	}

	return -1
}

func find(start string) (int, int) {
	var x, y int
	for i := 0; i < len(start); i++ {
		if start[i] == 'x' {
			x = i / 3
			y = i % 3
		}
	}

	return x, y
}
