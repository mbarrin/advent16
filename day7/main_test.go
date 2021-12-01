package main

import "testing"

func TestTotalValid(t *testing.T) {
	input := parseInput("test.txt")

	count := 2

	output := totalValid(input)

	if output != count {
		t.Logf("expected %d got %d", count, output)
		t.Fail()
	}
}
