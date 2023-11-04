package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func createFile(filePath string) (file *os.File) {
	var err error

	if file, err = os.Create(filePath); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	return
}

func openFile(filePath string) (file *os.File) {
	var err error

	if file, err = os.Open(filePath); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())

		os.Exit(1)
	}

	return
}

func parseFile(file *os.File) (tables tables) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		tokens := strings.Split(text, "\t")

		lastIdx := len(tables) - 1

		switch tokens[0] {
		case "":
			switch tokens[1] {
			case "FOREIGN KEY":
				tables[lastIdx].constraints.newForeignConstraint(tokens)
			case "PRIMARY KEY":
				tables[lastIdx].constraints.newPrimaryConstraint(tokens)
			case "UNIQUE KEY":
				tables[lastIdx].constraints.newUniqueConstraint(tokens)
			case "TABLE INDEX":
				tables[lastIdx].indexes.newIndex(tokens)
			default:
				tables[lastIdx].columns.newColumn(tokens)
			}
		default:
			tables.newTable(tokens)
		}
	}

	return
}

func saveToFile(outputFilepath, fileContent string) {
	var file *os.File

	if outputFilepath == "" {
		file = os.Stdout

	} else {
		file = createFile(outputFilepath)
		defer file.Close()
	}

	file.WriteString(fileContent)
}

func main() {
	flags := parseFlags()

	file := openFile(flags.filepath)
	defer file.Close()

	tables := parseFile(file)

	saveToFile(flags.outputFilepath, tables.script())
}
