package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func decryptV2(encrypted []string) int {
	length, endIndex := 0, 0
	var (
		count   int
		repeats int
	)

	fmt.Printf("encrypted: %v\n", encrypted)
	for i := 0; i < len(encrypted); i++ {
		fmt.Printf("length: %d\n", length)
		fmt.Printf("i: %d\n", i)
		if encrypted[i] == "(" {
			for j := i; j < len(encrypted); j++ {
				if encrypted[j] == ")" {
					endIndex = j
					break
				}
			}

			foo := strings.Join(encrypted[i:endIndex+1], "")
			//fmt.Printf("foo: %v\n", foo)
			fmt.Sscanf(foo, "(%dx%d)", &count, &repeats)

			tmp := decryptV2(encrypted[endIndex+1 : endIndex+1+count])
			toAdd := repeats * tmp
			//fmt.Printf("repeats: %v\n", repeats)
			fmt.Printf("tmp: %v\n", tmp)
			//fmt.Printf("toAdd: %v\n", toAdd)
			length += toAdd
			i += count + endIndex - 1
		} else {
			length++
		}
	}
	return length
}

func decrypt(split []string) int {
	outputLength, startIndex, endIndex := 0, 0, 0

	var (
		count   int
		repeats int
	)

	for i := 0; i < len(split); i++ {
		if split[i] == "(" {
			startIndex = i
			for j := i; j < len(split); j++ {
				if split[j] == ")" {
					endIndex = j
					break
				}
			}
			expander := strings.Join(split[startIndex:endIndex+1], "")

			fmt.Sscanf(expander, "(%dx%d)", &count, &repeats)

			outputLength += count * repeats
			i += len(expander) - 1 + count
		} else {
			outputLength++
		}
	}

	return outputLength
}

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		fmt.Printf("part 1: %d\n", decrypt(line))
		fmt.Printf("part 2: %d\n", decryptV2(line))
	}
}
