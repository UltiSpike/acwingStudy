package main

import "fmt"

const N = 5

var nums [N][N]int
var matrixSum [N][N]int

func main() {
	var n, m, q int
	_, _ = fmt.Scanf("%d %d %d", &n, &m, &q)

	for i := 1; i <= n; i++ {
		fmt.Scanf("\n")
		for j := 1; j <= m; j++ {
			_, _ = fmt.Scanf("%d", &nums[i][j])
			matrixSum[i][j] = matrixSum[i-1][j] + matrixSum[i][j-1] - matrixSum[i-1][j-1] + nums[i][j]
		}
	}

	//15 minutes
	var x1, x2, y1, y2 int
	for i := q; i > 0; i-- {
		fmt.Scanf("\n %d %d %d %d", &x1, &y1, &x2, &y2)
		fmt.Println(matrixSum[x2][y2] + matrixSum[x1-1][y1-1] - matrixSum[x2][y1-1] - matrixSum[x1-1][y2])
	}
}
