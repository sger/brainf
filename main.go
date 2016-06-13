package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const dataSize int = 65535

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}

	data, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Printf("Cannot open file %s\n", args[1])
		return
	}

	fmt.Println(string(data))
}
