package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func unescape(s string) string {
	if u, err := strconv.Unquote(`"` + s + `"`); err == nil {
		return u
	}
	r := strings.NewReplacer(`\t`, "\t", `\n`, "\n")
	return r.Replace(s)
}

func main() {
	del := flag.String("d", "\t", "delimeters to use")
	fields := flag.Int("f", 2, "fields to extract (1-based index)")
	flag.Parse()
	args := flag.Args()

	delimiter := unescape(*del)
	filename := "sample.tsv"
	if len(args) > 0 {
		filename = args[0]
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // Get the line as a string
		words := strings.Split(line, delimiter)
		for i, word := range words {
			if i == *fields-1 {
				fmt.Println(word)
			}
		}
	}
}
