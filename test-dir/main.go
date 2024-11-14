package main

import (
	"fmt"

	"github.com/aggre-golang/test-dir/util"
)

func main() {
	fmt.Println("main")

	row, col := 10, 7
	dp := util.CreateInfMatrix(row, col)
	dp[0][0] = 0
	fmt.Println(dp)
}
