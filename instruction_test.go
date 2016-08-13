package main

import "testing"

func TestOperator(t *testing.T) {
	expectedOperator := opIncDp
	instruction := Instruction{operator: opIncDp}

	if instruction.operator != expectedOperator {
		t.Fatalf("Expected %s, got %s", expectedOperator, instruction.operator)
	}
}

func TestOperand(t *testing.T) {
	expectedOperand := 0
	instruction := Instruction{operand: 0}

	if instruction.operand != expectedOperand {
		t.Fatalf("Expected %s, got %s", expectedOperand, instruction.operand)
	}
}
