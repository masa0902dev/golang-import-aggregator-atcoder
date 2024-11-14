package code

import "github.com/aggre-golang/test-dir/util"

func SomeProblem() [][]int {
	a := 10
	b := 20
	matrix := util.CreateInfMatrix(a, b)
	return matrix
}
