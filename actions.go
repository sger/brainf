package main

import (
	"fmt"
	"io/ioutil"

	"github.com/codegangsta/cli"
)

func parseFile(c *cli.Context) error {
	if len(c.Args().First()) > 0 {
		data, err := ioutil.ReadFile(c.Args().First())
		if err != nil {
			fmt.Printf("Cannot open file %s\n", c.Args().First())
			return nil
		}

		fmt.Println(string(data))
	}
	return nil
}
