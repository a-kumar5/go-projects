package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//This is go implementation for unix wc
	var tlc, twc, tcc int
	for _, fname := range os.Args[1:] {
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
		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
	if len(os.Args) > 2 {
		fmt.Printf("%7d %7d %7d %s\n", tlc, twc, tcc, "total")
	}
}
