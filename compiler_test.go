package main

import "testing"

func TestCompilerLengthOutput(t *testing.T) {

	var program = "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	instructions, err := compile(program)

	if err != nil {

	}

	if len(instructions) != len(program) {
		//t.Fatalf("Expected %s, got %s", programOutput, word)
	}

	err = output(instructions)

	if err != nil {

	}
}
