package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "greeting",
		Usage: "Prints greeting + name",
		Action: func(c *cli.Context) error {
			name := "world"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			fmt.Printf("hello, %s\n", name)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
