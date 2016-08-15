package main

func compile(instructions string) (chan<- rune, <-chan rune) {
	memory := make([]int, SIZE)
	var ptr = 0

	input := make(chan rune)
	output := make(chan rune)

	var counter = 0

	go func() {

		defer close(input)
		defer close(output)

		for counter < len(instructions) {
			switch instructions[counter] {
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
				if memory[ptr] == 0 {
					nestedCounter := 1
					for nestedCounter > 0 {
						counter++
						operator := instructions[counter]
						if operator == '[' {
							nestedCounter++
						} else if operator == ']' {
							nestedCounter--
						}
					}
					counter--
				}
			case ']':
				if memory[ptr] > 0 {
					nestedCounter := 1
					for nestedCounter > 0 {
						counter--
						operator := instructions[counter]
						if operator == ']' {
							nestedCounter++
						} else if operator == '[' {
							nestedCounter--
						}
					}
					counter--
				}
			default:

			}
			counter++
		}

	}()

	return input, output
}
