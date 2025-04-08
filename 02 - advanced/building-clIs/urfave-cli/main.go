package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// TODO
func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "Prints a greeting",
		Action: func(c *cli.Context) error {
			fmt.Println("hello world")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
