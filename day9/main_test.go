package main

import (
	"strings"
	"testing"
)

func TestDecrypt(t *testing.T) {
	input := map[string]int{
		"ADVENT":            6,
		"A(1x5)BC":          7,
		"(3x3)XYZ":          9,
		"A(2x2)BCD(2x2)EFG": 11,
		"(6x1)(1x3)A":       6,
		"X(8x2)(3x3)ABCY":   18,
	}

	for k, v := range input {
		decryptedLength := decrypt(strings.Split(k, ""))
		if decryptedLength != v {
			t.Fatalf("expected %d, got %d", v, decryptedLength)
		}
	}
}

func TestDecryptV2(t *testing.T) {
	input := map[string]int{
		//"(3x3)XYZ":                           9,
		"X(8x2)(3x3)ABCY": 20,
		//"(27x12)(20x12)(13x14)(7x10)(1x12)A":                       241920,
		//"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN": 445,
	}

	for k, v := range input {
		decryptedLength := decryptV2(strings.Split(k, ""))
		if decryptedLength != v {
			t.Fail()
			t.Logf("expected %d, got %d", v, decryptedLength)
		}
	}
}
