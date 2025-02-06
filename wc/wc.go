package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
func getAllFlags(args []string) []string {
	if len(args) > 1{

	}
}
*/

type fMetadata struct {
	fname     string
	LineCount int
	WordCount int
	CharCount int
}

func getFileMetadata(fname string) fMetadata {
	var fm fMetadata
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fm.fname = fname
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		fm.WordCount += len(strings.Fields(s))
		fm.CharCount += len(s)
		fm.LineCount++
	}
	return fm
}

func main() {
	//This is go implementation for unix wc

	// get flags from the users.
	var printLineCount, printWordCount, printCharCount bool
	flag.BoolVar(&printLineCount, "l", false, "to only get line count")
	flag.BoolVar(&printWordCount, "w", false, "to only get word count")
	flag.BoolVar(&printCharCount, "c", false, "to only get character count")
	flag.Parse()
	fnames := flag.CommandLine.Args()

	// variables defined to store total line count, word count and character count
	var total fMetadata
	for _, fname := range fnames {
		fm := getFileMetadata(fname)
		fmt.Printf("%7d %7d %7d %s\n", fm.LineCount, fm.WordCount, fm.CharCount, fm.fname)
		total.fname = "total"
		total.LineCount += fm.LineCount
		total.WordCount += fm.WordCount
		total.CharCount += fm.CharCount
	}
	if len(fnames) > 1 {
		if printLineCount {
			fmt.Printf("%7d %s\n", total.LineCount, "total")
		} else if printCharCount {
			fmt.Printf("%7d %s\n", total.CharCount, "total")
		} else if printWordCount {
			fmt.Printf("%7d %s\n", total.WordCount, "total")
		} else {
			fmt.Printf("%7d %7d %7d %s\n", total.LineCount, total.WordCount, total.CharCount, "total")
		}
	}
}
