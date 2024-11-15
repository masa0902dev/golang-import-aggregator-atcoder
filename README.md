# 🔥golang-import-aggregator

> [日本語 README はこちら](https://github.com/masa0902dev/golang-import-aggregator-atcoder/blob/main/README-ja.md)

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Demo](#demo)
  - [Before Aggregation (3 Files)](#before-aggregation-3-files)
  - [After Aggregation](#%EF%B8%8Fafter-aggregation)
- [Prerequisites](#prerequisites)
- [License](#license)

## Features

This tool consolidates custom packages imported in a main file into a single main file.  
It is especially useful when using Golang for competitive programming, such as on platforms like AtCoder.

The result is output to the console, and you can choose to overwrite the specified main file (you can skip confirmation).

## Installation

```bash
go install github.com/masa0902dev/golang-import-aggregator-atcoder/cmd/agg
```

<br>

## 🔥Usage

To consolidate imports into a main file, use the following command.  
Note: Imports are not automatically added, so please use your IDE's linting tool to include them!

[Example]

```bash
agg \
-main test-dir/chap1/main.go \
-import test-dir/chap1/code/problem.go,test-dir/util/util.go \
-prefix code,util \
-skip
```

Specify paths as relative paths from the current directory.

- `-main` specifies the path to the main file.

- `-import` specifies the paths of packages imported by the main file, packages imported within those packages, and so on.

- `-prefix` specifies the package names used.

- `-skip` if you add skip flag, you can paste the code to Main file without confirmation in terminal.

<br>

In the example above, the directory structure is as follows.  
The import relationship is that main.go imports the code package, and code/problem.go imports the util package.

```bash
test-dir
│
├── chap1
│   ├── code
│   │   └── problem.go
│   └── main.go
└── util
    └── util.go
```

<br>

## 🔥Demo

### Before Aggregation (3 Files)

```go
// chap1/main.go ------------------------------------------
package main

import (
	"fmt"

	"github.com/masa0902dev/golang-import-aggregator-atcoder/test-dir/chap1/code"
)

func main() {
	fmt.Println(code.SomeProblem())
}


// chap1/code/problem.go ------------------------------------------
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


// util/util.go ------------------------------------------
package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CreateInfMatrix(row int, col int) [][]int {
	const inf = int(^uint(0) >> 1)
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
```

### ⚡️After Aggregation

Note: Imports are not automatically added, so please use your IDE's linting tool to include them!

```go
// chap1/main.go ------------------------------------------
package main


func main() {
	fmt.Println(Problem())
}

func CreateInfMatrix(row int, col int) [][]int {
	const inf = int(^uint(0) >> 1)
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

func Problem() [][]int {
	a := 10
	b := 20
	matrix := CreateInfMatrix(a, b)
	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] -= 1
		}
	}
	return matrix
}
```

<br><br>

## Prerequisites

Go 1.16 or later

## License

This project is licensed under the MIT License.  
See the LICENSE file for details.
