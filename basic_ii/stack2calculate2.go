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
	line = strings.TrimSpace(line) // 返回的是一个 string字符串
	nums := make([]int, 0)
	op := make([]byte, 0)
	table := map[byte]int{
		'*': 2,
		'/': 2,
		'+': 1,
		'-': 1,
	}

	for i := 0; i < len(line); i++ {
		// 		fmt.Println(line[i])
		if line[i] >= '0' && line[i] <= '9' {
			j := i
			x := 0
			for j < len(line) && line[j] >= '0' && line[j] <= '9' {
				x = x*10 + int(line[j]-'0')
				j++
			}
			i = j - 1
			nums = append(nums, x)
		} else {
			if line[i] == '(' {
				op = append(op, '(')
			} else {
				if line[i] == ')' {
					for op[len(op)-1] != '(' {
						cal(&op, &nums)
					}
					op = op[:len(op)-1]
				} else {
					// 弹栈要求栈顶元素优先级比较高 不能等后面的结果
					for len(op) > 0 && table[line[i]] <= table[op[len(op)-1]] {
						cal(&op, &nums)
					}
					op = append(op, line[i])
				}
			}
		}
	}
	for len(op) > 0 {
		cal(&op, &nums)
	}
	fmt.Println(nums[0])
}

// 引用传递
func cal(op *[]byte, nums *[]int) {
	a := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	b := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	c := (*op)[len(*op)-1]
	*op = (*op)[:len(*op)-1]

	var d int
	switch c {
	case '*':
		d = a * b
	case '/':
		d = b / a
	case '+':
		d = b + a
	case '-':
		d = b - a
	}

	*nums = append(*nums, d)
}
