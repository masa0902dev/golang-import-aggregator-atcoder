package main

import "fmt"

func main() {
	fmt.Println("main")

	row, col := 10, 7
	dp := CreateInfMatrix(row, col)
	dp[0][0] = 0
	fmt.Println(dp)
}

func CreateInfMatrix(row int, col int) [][]int {
	const inf = int(^uint(0) >> 1) // maximum int value
	matrix := make([][]int, row)
	for i := range matrix {
		matrix[i] = make([]int, col)
		for j := range matrix[i] {
			matrix[i][j] = inf
		}
	}
	return matrix
}
