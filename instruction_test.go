package main

import "testing"

func TestOperator(t *testing.T) {
	expectedOperator := opIncDp
	instruction := Instruction{operator: opIncDp}

	if instruction.operator != expectedOperator {
		t.Fatalf("Expected %d, got %d", expectedOperator, instruction.operator)
	}
}

func TestOperand(t *testing.T) {
	expectedOperand := 0
	instruction := Instruction{operand: 0}

	if instruction.operand != expectedOperand {
		t.Fatalf("Expected %d, got %d", expectedOperand, instruction.operand)
	}
}
