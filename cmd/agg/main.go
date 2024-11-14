package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/masa0902dev/golang-import-aggregator-atcoder/pkg"
)

func main() {
	mainPath := flag.String("main", "", "Main Go file path")
	importPaths := flag.String("import", "", "Comma-separated list of import file paths")
	prefixesToRemove := flag.String("prefix", "", "Comma-separated list of prefixes to remove")
	flag.Parse()

	if *mainPath == "" || *importPaths == "" || *prefixesToRemove == "" {
		fmt.Println("Usage: aggre -main <main file path> -import <comma-separated import paths> -prefix <comma-separated prefixes>")
		os.Exit(1)
	}

	importFiles := strings.Split(*importPaths, ",")
	prefixes := strings.Split(*prefixesToRemove, ",")

	finalCode, err := pkg.AggregateAndProcessFiles(*mainPath, importFiles, prefixes)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(finalCode)

	if pkg.GetUserConfirmation("Do you want to overwrite the main file with this code? (y/n): ") {
		if err := pkg.OverwriteFile(*mainPath, finalCode); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("File overwritten successfully.")
	} else {
		fmt.Println("File not overwritten.")
	}
}
