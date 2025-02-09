package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type fMetadata struct {
	lines int
	words int
	chars int
}

func getFileMetadata(fname string) (fMetadata, error) {
	file, err := os.Open(fname)
	if err != nil {
		return fMetadata{}, err
	}
	defer file.Close()

	var meta fMetadata
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		s := scan.Text()
		meta.words += len(strings.Fields(s))
		meta.chars += len(s)
		meta.lines++
	}
	return meta, scan.Err()
}

func printMetadata(fname string, meta fMetadata, flags map[string]bool) {
	switch {
	case flags["l"]:
		fmt.Printf("%7d %s\n", meta.lines, fname)
	case flags["w"]:
		fmt.Printf("%7d %s\n", meta.words, fname)
	case flags["c"]:
		fmt.Printf("%7d %s\n", meta.chars, fname)
	default:
		fmt.Printf("%7d %7d %7d %s\n", meta.lines, meta.words, meta.chars, fname)
	}
}

func main() {
	var l, w, c bool
	flag.BoolVar(&l, "l", false, "line count")
	flag.BoolVar(&w, "w", false, "word count")
	flag.BoolVar(&c, "c", false, "character count")
	flag.Parse()
	fnames := flag.CommandLine.Args()

	flags := map[string]bool{
		"l": l,
		"w": w,
		"c": c,
	}

	// variables defined to store total line count, word count and character count
	var total fMetadata
	for _, fname := range fnames {
		meta, err := getFileMetadata(fname)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			continue
		}
		printMetadata(fname, meta, flags)
		total.lines += meta.lines
		total.words += meta.words
		total.chars += meta.chars
	}
	if len(fnames) > 1 {
		printMetadata("total", total, flags)
	}
}
