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
	var tlc, twc, tcc int

	for _, fname := range fnames {
		var lc, wc, cc int
		file, err := os.Open(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		tlc += lc
		twc += wc
		tcc += cc
		if printLineCount {
			fmt.Printf("%7d %s\n", lc, fname)
			file.Close()
		} else if printWordCount {
			fmt.Printf("%7d %s\n", wc, fname)
			file.Close()
		} else if printCharCount {
			fmt.Printf("%7d %s\n", cc, fname)
			file.Close()
		} else {
			fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
			file.Close()
		}
	}
	if len(fnames) > 1 {
		if printLineCount {
			fmt.Printf("%7d %s\n", tlc, "total")
		} else if printCharCount {
			fmt.Printf("%7d %s\n", tcc, "total")
		} else if printWordCount {
			fmt.Printf("%7d %s\n", twc, "total")
		} else {
			fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, "total")
		}
	}
}
