# 🔥golang-import-aggregator

## 目次

- [できること](#できること)
- [インストール](#インストール)
- [使い方](#使い方)
- [動作イメージ](#動作イメージ)
  - [統合前 (３ファイル)](#統合前-３ファイル)
  - [統合後](#%EF%B8%8F統合後)
- [前提条件](#前提条件)
- [ライセンス](#ライセンス)

## できること

Main ファイルにおいて import された自作パッケージ群を、Main ファイルにひとまとめにできます。  
特に AtCoder などの競技プログラミングで Golang を使用する場合に役立ちます。

コンソールに結果が出力され、指定の Main ファイルへ上書きするかを選択できます(確認のスキップ可能)。

## インストール

```bash
go install github.com/masa0902dev/golang-import-aggregator-atcoder/cmd/agg
```

<br>

## 🔥 使い方

import を Main ファイルにひとまとめにするには以下のように使用します。  
<b>ただし、import は記述されないので、自身の IDE の lint などで import を追記して下さい!</b>

[使用例]

```bash
agg \
-main test-dir/chap1/main.go \
-import test-dir/chap1/code/problem.go,test-dir/util/util.go \
-prefix code,util \
-skip
```

パスはカレントディレクトリからの相対パスで指定します。

- `-main` Main ファイルのパスを指定

- `-import` Main ファイルで import しているパッケージのパス、そのパッケージで import しているパッケージのパス、...を指定

- `-prefix` 使用しているパッケージ名を指定

- `-skip` skip フラッグをつけると、ターミナル上での確認無しで Main ファイルにコードをペーストできます

<br>

上例では、以下のようなディレクトリ構造になっています。  
import 関係は、main.go が code パッケージを import, code/problem.go が util パッケージを import しています。

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

## 🔥 動作イメージ

### 統合前 (３ファイル)

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

### ⚡️ 統合後

<b>import は記述されないので、自身の IDE の lint などで import を追記して下さい!</b>

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

<br>
<br>

## 前提条件

- Go 1.16 以降

## ライセンス

このプロジェクトは MIT ライセンスの下でライセンスされています。  
詳細は[LICENSE](LICENSE)ファイルを参照してください。
