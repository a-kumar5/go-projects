package main

import (
	"fmt"
	"os"

	hello "github.com/a-kumar5/matt-course"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
