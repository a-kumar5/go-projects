package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var KnownCommands = map[string]int{"exit": 0, "echo": 1, "type": 2}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		input = strings.TrimRight(input, "\n")
		tokenizedInput := strings.Split(input, " ")
		cmd := tokenizedInput[0]
		if fn, exists := KnownCommands[cmd]; !exists {
			if output := RunCmd(tokenizedInput); len(output) > 0 {
				fmt.Printf("%s", output)
			} else {
				fmt.Fprintf(os.Stdout, "%v: command not found\n", input)
			}
		} else {
			switch fn {
			case 0:
				DoExit(tokenizedInput[1:])
			case 1:
				DoEcho(tokenizedInput[1:])
			case 2:
				DoType(tokenizedInput[1:])
			}
		}
	}
}

func RunCmd(args []string) []byte {
	cmd := exec.Command(args[0], args[1:]...)
	b, _ := cmd.CombinedOutput()

	return b
}

func DoType(params []string) {
	item := params[0]
	if _, exists := KnownCommands[item]; exists {
		fmt.Fprintf(os.Stdout, "%v is a shell builtin\n", item)
	} else {
		env := os.Getenv("PATH")
		paths := strings.Split(env, ":")
		for _, path := range paths {
			exec := path + "/" + item
			if _, err := os.Stat(exec); err == nil {
				fmt.Fprintf(os.Stdout, "%v is %v\n", item, exec)
				return
			}
		}
		fmt.Fprintf(os.Stdout, "%v: not found\n", item)
	}
}

func DoEcho(params []string) {
	output := strings.Join(params, " ")
	fmt.Fprintf(os.Stdout, "%v\n", output)
}

func DoExit(params []string) {
	os.Exit(0)
}

func printPrompt() {
	fmt.Fprint(os.Stdout, "$ ")
}
