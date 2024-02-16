package basic_ii_exe

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const I = 100010

var p [I]int
var d [I]int

func union(x, y, ret int) {
	fx := find(x)
	fy := find(y)

	p[fx] = fy
	d[fx] = ret

}

func find(x int) int {
	if x != p[x] {
		temp := find(p[x])
		d[x] += d[p[x]]
		p[x] = temp
	}

	return p[x]
}

func main() {
	var n, m int
	_, _ = fmt.Scanln(&n, &m)

	reader := bufio.NewReader(os.Stdin)

	for i := 1; i <= n; i++ {
		d[i] = 0
		p[i] = i
	}

	var ans int
	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")

		ret, _ := strconv.Atoi(words[0])
		x, _ := strconv.Atoi(words[1])
		y, _ := strconv.Atoi(words[2])
		if x > n || y > n {
			ans++
			continue
		}
		fx := find(x)
		fy := find(y)
		if fx != fy {
			union(x, y, ret-1)
		} else {
			if ((d[fx]-d[fy])%3+3)%3 != ret-1 {
				ans += 1
			}
		}

	}

	fmt.Println(ans)
}
