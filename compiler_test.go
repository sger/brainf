package main

import (
	"fmt"
	"testing"
)

func TestCompilerOutput(t *testing.T) {

	var programOutput = "Hello World!"

	instructions, err := compile("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")

	if err != nil {

	}

	word := output(instructions)
	fmt.Println(word)

	if word != programOutput {
		t.Fatalf("Expected %s, got %s", programOutput, word)
	}
}
