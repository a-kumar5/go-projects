package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	//command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//fmt.Fprintf(os.Stdout, command[:len(command)-1]+": command not found\n")
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		/*
			text := reader.Text()
			if text := reader.Text(); text == "exit 0" {
				os.Exit(0)
			}
			fmt.Fprintf(os.Stdout, strings.TrimSpace(text)+": command not found\n")
			printPrompt()
		*/
		cmdline := strings.Fields(strings.TrimSpace(reader.Text()))
		cmd := cmdline[0]
		args := strings.Join(cmdline[1:], " ")
		switch cmd {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Fprintf(os.Stdout, args+"\n")
		default:
			fmt.Fprintf(os.Stdout, strings.Join(cmdline, " ")+": command not found\n")
		}
		printPrompt()
	}
}

func printPrompt() {
	fmt.Fprint(os.Stdout, "$ ")
}
