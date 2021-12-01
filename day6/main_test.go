package main

import "testing"

func TestMessage(t *testing.T) {
	input := parseInput("test.txt")
	partOneExpected := "easter"
	partTwoExpected := "advent"

	partOne := message(input, "most")
	partTwo := message(input, "least")

	if partOne != partOneExpected {
		t.Logf("decoded should be %s, but got %s", partOneExpected, partOne)
		t.Fail()
	}

	if partTwo != partTwoExpected {
		t.Logf("decoded should be %s, but got %s", partTwoExpected, partTwo)
		t.Fail()
	}
}
