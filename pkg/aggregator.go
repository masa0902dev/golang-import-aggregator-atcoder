package pkg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"regexp"
	"strings"
)

func parseFileToASTString(filePath string, funcMap map[string]bool) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	var output strings.Builder
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if _, exists := funcMap[funcDecl.Name.Name]; exists {
				continue
			}
			funcMap[funcDecl.Name.Name] = true
		}
		if err := printer.Fprint(&output, fset, decl); err != nil {
			return "", err
		}
		output.WriteString("\n")
	}

	return output.String(), nil
}

func removeUnwantedLines(code string, prefixes []string) string {
	code = regexp.MustCompile(`(?m)^package\s+\w+\s*$`).ReplaceAllString(code, "")
	code = regexp.MustCompile(`(?m)^\s*import\s+\([\s\S]*?\)\s*$|^\s*import\s+"[\w./-]+".*?$`).ReplaceAllString(code, "")

	for _, prefix := range prefixes {
		code = regexp.MustCompile(fmt.Sprintf(`\b%s\.`, regexp.QuoteMeta(prefix))).ReplaceAllString(code, "")
	}

	code = regexp.MustCompile(`\n{2,}`).ReplaceAllString(code, "\n\n")

	return code
}

func AggregateAndProcessFiles(mainFile string, importFiles []string, prefixes []string) (string, error) {
	combinedCode := ""
	funcMap := make(map[string]bool)

	mainCode, err := parseFileToASTString(mainFile, funcMap)
	if err != nil {
		return "", err
	}
	combinedCode += mainCode + "\n"

	for _, filePath := range importFiles {
		importCode, err := parseFileToASTString(filePath, funcMap)
		if err != nil {
			return "", err
		}
		combinedCode += importCode + "\n"
	}

	finalCode := removeUnwantedLines(combinedCode, prefixes)

	return "package main\n\n" + finalCode, nil
}

func OverwriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}
