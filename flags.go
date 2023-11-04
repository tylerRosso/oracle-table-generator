package main

import "flag"

type flags struct {
	filepath       string
	outputFilepath string
}

func parseFlags() (flags flags) {
	flag.StringVar(&flags.filepath, "filepath", "", "path to the file containing the table definition")
	flag.StringVar(&flags.outputFilepath, "output-filepath", "", "path to the resulting file (the one containing the table script)")

	flag.Parse()

	return
}
