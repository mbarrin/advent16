package main

import "testing"

func TestTotalValid(t *testing.T) {
	input := parseInput("test.txt")

	count := 2

	tls, _ := totalValid(input)

	if tls != count {
		t.Logf("expected %d got %d", count, tls)
		t.Fail()
	}
}
