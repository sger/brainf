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

				parseFile(c)
				return nil
			},
		},
		{
			Name:    "string",
			Aliases: []string{"s"},
			Usage:   "./brainf s `++++++++++`",
			Action: func(c *cli.Context) error {

				parseString(c)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
