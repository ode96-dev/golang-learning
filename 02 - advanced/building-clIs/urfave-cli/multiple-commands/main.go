package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "toolkit",
		Usage: "a simple cli with multiple commands",
		Commands: []*cli.Command{
			{
				Name:  "greet",
				Usage: "print greeting",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Value: "world"},
				},
				Action: func(ctx *cli.Context) error {
					fmt.Printf("hello, %s\n", ctx.String("name"))
					return nil
				},
			},
			{
				Name:  "bye",
				Usage: "say goodbye",
				Action: func(ctx *cli.Context) error {
					fmt.Printf("goodbye")
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
