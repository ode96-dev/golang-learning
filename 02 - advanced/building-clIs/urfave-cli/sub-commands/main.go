package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "admin",
		Usage: "Admin Commands",
		Commands: []*cli.Command{
			{
				Name:  "user",
				Usage: "Manage users",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new user",
						Action: func(ctx *cli.Context) error {
							fmt.Println("user added")
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete a user",
						Action: func(ctx *cli.Context) error {
							fmt.Println("user deleted")
							return nil
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
