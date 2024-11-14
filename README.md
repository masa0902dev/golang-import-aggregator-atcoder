# golang-import-aggregator

## できること

Mainファイルにおいてimportされた自作パッケージ群を、Mainファイルにひとまとめにできます。  
特にAtCoderなどの競技プログラミングでGolangを使用する場合に役立ちます。

コンソールに結果が出力され、指定のMainファイルへの上書きを選択できます。



## インストール
```bash
go install github.com/masa0902dev/golang-import-aggregator-atcoder
```


<br>

## 使い方

importをmain.goにひとまとめにするには以下のように使用します。  
<b>ただし、importは記述されないので、自身のIDEのlintなどでimportを追記して下さい!</b>

[使用例]
```bash
agg \
-main test-dir/chap1/main.go \
-import test-dir/chap1/code/problem.go,test-dir/util/util.go \
-prefix code,util
```

パスはカレントディレクトリからの相対パスで指定します。


- `-main` Mainファイルのパスを指定

- `-import` Mainファイルでimportしているパッケージのパス、そのパッケージでimportしているパッケージのパス、...を指定

- `-prefix` 使用しているパッケージ名を指定

<br>

上例では、以下のようなディレクトリ構造になっています。  
import関係は、main.goがcodeパッケージをimport, code/problem.goがutilパッケージをimportしています。
```bash
.
├── chap1
│   ├── code
│   │   └── problem.go
│   └── main.go
└── util
    └── util.go
```



## 動作イメージ
上例のコマンドを実行した場合の動作イメージは、以下の通りです。


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

### 🔥統合後🔥
<b>importは記述されないので、自身のIDEのlintなどでimportを追記して下さい!</b>
```go
package main


func main() {
	fmt.Println(Problem())
}

func CreateInfMatrix(row int, col int) [][]int {
	const inf = int(^uint(0) >> 1)	// maximum int value
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

## 前提条件

- Go 1.16以降



## ライセンス

このプロジェクトはMITライセンスの下でライセンスされています。  
詳細は[LICENSE](LICENSE)ファイルを参照してください。
