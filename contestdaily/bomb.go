package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const M = 5010

func main() {
	var n, r int
	_, _ = fmt.Scanln(&n, &r)
	if r == 0 {
		fmt.Println("0")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var s [M][M]int
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		words := strings.Split(line, " ")
		x, _ := strconv.Atoi(words[0])
		y, _ := strconv.Atoi(words[1])
		w, _ := strconv.Atoi(words[2])
		s[x+1][y+1] += w // 如果只需要前缀和数组 可以使用这种方式
	}

	for i := 1; i <= 5001; i++ {
		cnt := 0
		for j := 1; j <= 5001; j++ {
			cnt += s[i][j]
			s[i][j] = s[i-1][j] + cnt
		}
	}

	var ans int
	// 二位前缀和  area(x1,y1,x2,y2) = per(x2,y2) - pre(x2,y1-1)-pre(x1,y2-1) + pre(x1-1,y1-1)
	for i := 1; i <= 5001; i++ {
		for j := 1; j <= 5001; j++ {
			x2, y2 := i+r-1, j+r-1
			if x2 > 5001 {
				x2 = 5001
			}
			if y2 > 5001 {
				y2 = 5001
			}
			area := s[x2][y2] - s[i-1][y2] - s[x2][j-1] + s[i-1][j-1]
			if area > ans {
				ans = area
			}
		}
	}

	if r > 5001 {
		r = 5001
	}
	//var ans int
	// 更好的写法
	// r 在计算格子时需要-1 所以实际x1，y1也自动-1是边界的上一格了
	//for i := r; i <= 5001; i++ {
	//	for j := r; j <= 5001; j++ {
	//		area := s[i][j] - s[i-r][j] - s[i][j-r] + s[i-r][j-r]
	//		if area > ans {
	//			ans = area
	//		}
	//	}
	//}

	fmt.Println(ans)

}

//
