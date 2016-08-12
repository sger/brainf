package main

import (
	"fmt"
	"io/ioutil"
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
			Usage:   "open a file with extension .bf",
			Action: func(c *cli.Context) error {

				if len(c.Args().First()) > 0 {
					data, err := ioutil.ReadFile(c.Args().First())
					if err != nil {
						fmt.Printf("Cannot open file %s\n", c.Args().First())
						return nil
					}

					fmt.Println(string(data))
				}
				return nil
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add",
			Action: func(c *cli.Context) error {

				return nil
			},
		},
	}

	app.Run(os.Args)
}
