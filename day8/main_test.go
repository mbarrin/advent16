package main

import "testing"

func TestRotateRow(t *testing.T) {

	input := []bool{true, false, true, false}

	amount := 3

	expectedOuput := []bool{false, true, false, true}

	output := rotateRow(input, amount)

	if !identical(output, expectedOuput) {
		t.Fail()
		t.Logf("expected: %v, got: %v", expectedOuput, output)
	}

}

func TestRotateColumn(t *testing.T) {

	input := [][]bool{
		{true, false, true, false},
		{false, false, true, false},
		{false, false, true, false},
	}

	amount := 4

	expectedOutput := [][]bool{
		{false, false, true, false},
		{true, false, true, false},
		{false, false, true, false},
	}

	output := rotateColumn(input, 0, amount)

	for i := range output {
		if !identical(output[i], expectedOutput[i]) {
			t.Fail()
			t.Logf("expected: %v, got: %v", expectedOutput, output)
		}
	}
}

func TestTurnOn(t *testing.T) {
	input := [][]bool{
		{true, false, true, false},
		{false, false, false, false},
		{false, false, true, false},
	}

	expectedOutput := [][]bool{
		{true, true, true, false},
		{true, true, true, false},
		{true, true, true, false},
	}

	output := turnOn(input, 3, 3)

	for i := range output {
		if !identical(output[i], expectedOutput[i]) {
			t.Fail()
			t.Logf("expected: %v, got: %v", expectedOutput, output)
		}
	}

}

func TestLit(t *testing.T) {
	input := [][]bool{
		{true, false, true, false},
		{false, false, false, false},
		{false, false, true, false},
	}

	expectedOutput := 3

	output := lit(input)

	if expectedOutput != output {
		t.Fail()
		t.Logf("expected: %v, got: %v", expectedOutput, output)

	}

}

func identical(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
