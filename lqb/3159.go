package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	var b string
	var ans int
	for _, v := range line {
		// v == ?
		if string(v) == "?" {
			if b != "x" {
				ans += 1
			} else {
				b = "?"
			}
		} else { // v == "0" / "1"
			if string(v) == b || b == "?" {
				ans += 1
				// 下一个数从头匹配
				b = "x"
			} else {
				b = string(v)
			}
		}
	}

	fmt.Println(ans)
}
