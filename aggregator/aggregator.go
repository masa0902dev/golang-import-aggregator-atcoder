package main

import (
	"bufio"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"strings"
)

// ファイルをASTで解析して、重複のない関数を出力用に変換する
func parseFileToASTString(filePath string, funcMap map[string]bool) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// ASTを巡回し、重複する関数を出力しないようにする
	var output strings.Builder
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if _, exists := funcMap[funcDecl.Name.Name]; exists {
				continue // すでに出力されている関数はスキップ
			}
			funcMap[funcDecl.Name.Name] = true // 関数名を記録
		}
		// 出力に追加
		if err := printer.Fprint(&output, fset, decl); err != nil {
			return "", err
		}
		output.WriteString("\n")
	}

	return output.String(), nil
}

// package宣言や import文、指定プレフィックスを削除
func removeUnwantedLines(code string, prefixes []string) string {
	// package宣言を削除
	code = regexp.MustCompile(`(?m)^package\s+\w+\s*$`).ReplaceAllString(code, "")

	// すべてのimport文を削除（importブロックと単一行の両方に対応）
	code = regexp.MustCompile(`(?m)^\s*import\s+\([\s\S]*?\)\s*$|^\s*import\s+"[\w./-]+".*?$`).ReplaceAllString(code, "")

	// プレフィックスを削除
	for _, prefix := range prefixes {
		code = regexp.MustCompile(fmt.Sprintf(`\b%s\.`, regexp.QuoteMeta(prefix))).ReplaceAllString(code, "")
	}

	// 不要な空行を削除
	code = regexp.MustCompile(`\n{2,}`).ReplaceAllString(code, "\n\n")

	return code
}

// mainファイルとimportファイルを処理してプレフィックス削除
func aggregateAndProcessFiles(mainFile string, importFiles []string, prefixes []string) (string, error) {
	combinedCode := ""
	funcMap := make(map[string]bool) // 重複チェック用のマップ

	// mainファイルの内容を処理
	mainCode, err := parseFileToASTString(mainFile, funcMap)
	if err != nil {
		return "", err
	}
	combinedCode += mainCode + "\n"

	// importファイルの内容を処理
	for _, filePath := range importFiles {
		importCode, err := parseFileToASTString(filePath, funcMap)
		if err != nil {
			return "", err
		}
		combinedCode += importCode + "\n"
	}

	// 不要な行を削除
	finalCode := removeUnwantedLines(combinedCode, prefixes)

	return "package main\n\n" + finalCode, nil
}

func main() {
	mainPath := flag.String("main", "", "Main Go file path")
	importPaths := flag.String("import", "", "Comma-separated list of import file paths")
	prefixesToRemove := flag.String("prefixes", "", "Comma-separated list of prefixes to remove")
	flag.Parse()

	if *mainPath == "" || *importPaths == "" || *prefixesToRemove == "" {
		fmt.Println("Usage: aggre -main <main file path> -import <comma-separated import paths> -prefixes <comma-separated prefixes>")
		os.Exit(1)
	}

	// インポートファイルとプレフィックスをリスト化
	importFiles := strings.Split(*importPaths, ",")
	prefixes := strings.Split(*prefixesToRemove, ",")

	// プレフィックス削除後のコードを取得
	finalCode, err := aggregateAndProcessFiles(*mainPath, importFiles, prefixes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// 最終コードを出力
	fmt.Println(finalCode)

	// ユーザーに上書きの確認
	fmt.Print("Do you want to overwrite the main file with this code? (y/n): ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "y" {
			err = os.WriteFile(*mainPath, []byte(finalCode), 0644)
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("File overwritten successfully.")
		} else {
			fmt.Println("File not overwritten.")
		}
	}
}
