package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const T = 1000010

var son [T][26]int
var id int
var cnt [T]int

func insertT(str string) {
	var p int
	var u int
	for i := 0; i < len(str); i++ {
		u = int(str[i] - 'a')
		if son[p][u] == 0 {
			// new 节点
			id++
			son[p][u] = id
		}
		p = son[p][u]
	}
	cnt[p]++ // 以p为结尾的字符的单词个数+1
}

func queryT(str string) int {
	var p int
	var u int
	for i := 0; i < len(str); i++ {
		u = int(str[i] - 'a')
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}

	return cnt[p]
}

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")

		switch words[0] {
		case "I":
			insertT(words[1])
		case "Q":
			_, _ = writer.WriteString(strconv.Itoa(queryT(words[1])) + "\n")
		}

	}

	_ = writer.Flush()

}
