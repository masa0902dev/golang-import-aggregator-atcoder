package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func MultiMultiInt(maxBuffer int) [][]int {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 4096)
	scanner.Buffer(buffer, maxBuffer)
	var result [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		inputs := strings.Split(line, " ")
		nums := make([]int, len(inputs))
		for i, input := range inputs {
			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			nums[i] = num
		}
		result = append(result, nums)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil
	}

	return result
}
