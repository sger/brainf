package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/codegangsta/cli"
)

func parseFile(c *cli.Context) error {
	if len(c.Args().First()) > 0 {
		data, err := ioutil.ReadFile(c.Args().First())
		if err != nil {
			fmt.Printf("Cannot open file %s\n", c.Args().First())
			return nil
		}

		input, output := compile(string(data))

		go reader(input)
		printer(output)

	}
	return nil
}

func reader(input chan<- rune) {
	reader := bufio.NewReader(os.Stdin)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err.Error())
		}
		input <- r
	}
}

func printer(output <-chan rune) {
	for {
		output, ok := <-output
		if ok {
			fmt.Print(string(output))
		} else {
			break
		}
	}
}

func parseString(c *cli.Context) error {
	if len(c.Args().First()) > 0 {

		fmt.Println(c.Args().First())

		input, output := compile(c.Args().First())

		go reader(input)
		printer(output)
	}
	return nil
}
