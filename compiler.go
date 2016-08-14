package main

import "fmt"

func compile(instructions string) (chan<- rune, <-chan rune) {
	memory := make([]int, SIZE)
	var ptr = 0

	input := make(chan rune)
	output := make(chan rune)

	go func() {

		defer close(input)
		defer close(output)

		for i := 0; i < len(instructions); i++ {
			switch instructions[i] {
			case '>':
				ptr++
			case '<':
				ptr--
			case '+':
				memory[ptr]++
			case '-':
				memory[ptr]--
			case ',':
				memory[ptr] = int(<-input)
			case '.':
				output <- rune(memory[ptr])
			case '[':
				fmt.Println(memory[ptr])
				if memory[ptr] == 0 {

				}
			case ']':
				fmt.Println(memory[ptr])
				if memory[ptr] > 0 {

				}
			default:

			}
		}

	}()

	return input, output
}
