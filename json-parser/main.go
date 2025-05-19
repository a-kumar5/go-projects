package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Json Parsing Started")
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot read file: %v", err)
	}
	fmt.Printf("Contents of file is %s", string(data))

	if len(data) == 0 {
		os.Exit(1)
	}
}
