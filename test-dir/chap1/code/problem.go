package code

import "github.com/masa0902dev/golang-import-aggregator-atcoder/test-dir/util"

func Problem() [][]int {
	a := 10
	b := 20
	matrix := util.CreateInfMatrix(a, b)
	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] -= 1
		}
	}
	return matrix
}
