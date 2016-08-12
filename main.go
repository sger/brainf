package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	/*args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}

	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Printf("Cannot open file %s\n", args[1])
		return
	}

	fmt.Println(string(data))*/

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

				if len(c.Args().First()) > 0 {
					fmt.Println(c.Args().First())
				}

				return nil
			},
		},
	}

	app.Run(os.Args)
}
