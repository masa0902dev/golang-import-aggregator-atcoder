package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

// mainファイルとimportファイルを一つのファイルにまとめる
func main() {
	mainPath := flag.String("main", "", "Main Go file path")
	importPaths := flag.String("import", "", "Comma-separated list of import file paths")
	flag.Parse()

	if *mainPath == "" || *importPaths == "" {
		fmt.Println("Usage: aggre -main <main file path> -import <comma-separated import paths>")
		os.Exit(1)
	}

	// ファイルセットと出力ファイル用のASTファイルを生成
	fset := token.NewFileSet()
	output := &ast.File{Name: ast.NewIdent("main")}

	// mainファイルを解析し、出力に追加
	if err := parseFile(fset, *mainPath, output); err != nil {
		fmt.Printf("Error parsing main file: %v\n", err)
		os.Exit(1)
	}

	// importパスを処理し、各ファイルを解析して追加
	paths := strings.Split(*importPaths, ",")
	for _, path := range paths {
		if err := parseFile(fset, path, output); err != nil {
			fmt.Printf("Error parsing import file %s: %v\n", path, err)
			os.Exit(1)
		}
	}

	// まとめたASTをコマンドラインに出力
	if err := printer.Fprint(os.Stdout, fset, output); err != nil {
		fmt.Printf("Error printing combined file: %v\n", err)
	}
}

// 指定したパスのGoファイルをパースし、ASTに追加
func parseFile(fset *token.FileSet, path string, output *ast.File) error {
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// import文を削除して関数や型のみ追加
	for _, decl := range file.Decls {
		output.Decls = append(output.Decls, decl)
	}

	return nil
}
