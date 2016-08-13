package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func compile(tokens string) (instructions []Instruction, err error) {
	var programCounter = 0
	var stackProgramCounter = 0
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case '>':
			instructions = append(instructions, Instruction{operator: opIncDp})
		case '<':
			instructions = append(instructions, Instruction{operator: opDecDp})
		case '+':
			instructions = append(instructions, Instruction{operator: opIncVal})
		case '-':
			instructions = append(instructions, Instruction{operator: opDecVal})
		case ',':
			instructions = append(instructions, Instruction{operator: opIn})
		case '.':
			instructions = append(instructions, Instruction{operator: opOut})
		case '[':
			instructions = append(instructions, Instruction{operator: opJmpFwd})
			stack = append(stack, programCounter)
		case ']':
			stackProgramCounter = stack[len(stack)-1]
			instructions = append(instructions, Instruction{opJmpBck, stackProgramCounter})
			instructions[stackProgramCounter].operand = programCounter
		default:
			programCounter--
		}
		programCounter++
	}

	return instructions, errors.New("Error")
}

func output(instructions []Instruction) (err error) {
	data := make([]int16, SIZE)
	var ptr int = 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(instructions); i++ {
		switch instructions[i].operator {
		case opIncDp:
			ptr++
		case opDecDp:
			ptr--
		case opIncVal:
			data[ptr]++
		case opDecVal:
			data[ptr]--
		case opOut:
			fmt.Printf("%c", data[ptr])
		case opIn:
			value, _ := reader.ReadByte()
			data[ptr] = int16(value)
		case opJmpFwd:
			if data[ptr] == 0 {
				i = int(instructions[i].operand)
			}
		case opJmpBck:
			if data[ptr] > 0 {
				i = int(instructions[i].operand)
			}
		default:
			return errors.New("Error")
		}
	}
	return nil
}
