package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//		中缀表达式转后缀求值
//		双栈问题
//	 运算符栈
//	 匹配规则

const sb = 10000

var operate [sb]byte
var oo = -1
var nums [sb]int
var pp = -1

func main() {
	table := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	for i := 0; i < len(line); i++ {
		// 处理数字
		if line[i] >= '0' && line[i] <= '9' {
			x := 0
			j := i
			for j < len(line) && line[j] >= '0' && line[j] <= '9' {
				x = x*10 + int(line[j]-'0')
				j++
			}
			i = j - 1
			pp++
			nums[pp] = x
		} else { // 中缀转后缀 并且进行压栈
			if line[i] == '(' {
				oo++
				operate[oo] = line[i]
			} else {
				if line[i] == ')' {
					for operate[oo] != '(' {
						cal()
					}
					// 左括号也要出栈
					oo--
				} else {
					for oo > -1 && table[line[i]] <= table[operate[oo]] {
						cal()
					}
					oo++
					operate[oo] = line[i]
				}
			}
		}
		// fmt.Println(byte(line[i]), nums[:pp+1])
	}
	for oo > -1 {
		cal()
	}
	fmt.Println(nums[0])
}

// 计算结果并压栈
func cal() {
	a := nums[pp]
	pp--
	b := nums[pp]
	pp--
	op := operate[oo]
	oo--
	switch op {
	case '*':
		pp++
		nums[pp] = a * b
	case '+':
		pp++
		nums[pp] = a + b
	case '-':
		pp++
		nums[pp] = b - a
	case '/':
		pp++
		nums[pp] = b / a
	}

}
