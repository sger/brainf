package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Brainf"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "./brainf f hello_world.bf",
			Action: func(c *cli.Context) error {

				return nil
			},
		},
		{
			Name:    "string",
			Aliases: []string{"i"},
			Usage:   "./brainf i",
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}

	app.Run(os.Args)
}
