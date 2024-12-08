package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sesh",
		Usage: "Smart session manager for the terminal",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language of the greeting",
			},
		},
		Action: func(ctx *cli.Context) error {
			name := "Nefertiti"
			if ctx.NArg() > 0 {
				name = ctx.Args().Get(0)
			}
			if ctx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
