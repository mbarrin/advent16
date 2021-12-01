package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func transpose(s [][]string) [][]string {
	height := len(s)
	width := len(s[0])
	newSlice := make([][]string, width)

	for i := 0; i < width; i++ {
		newSlice[i] = make([]string, height)
		for j := 0; j < height; j++ {
			newSlice[i][j] = s[j][i]
		}
	}
	return newSlice
}

func mostCommon(s []string) string {
	var mostCommonKey string

	count := make(map[string]int)

	for _, x := range s {
		count[x]++
	}

	for k, v := range count {
		if v > count[mostCommonKey] {
			mostCommonKey = k
		}
	}

	return mostCommonKey
}

func leastCommon(s []string) string {
	var leastCommonKey string

	count := make(map[string]int)

	for _, x := range s {
		count[x]++
	}

	for k, v := range count {
		if count[leastCommonKey] == 0 {
			leastCommonKey = k
		} else if v < count[leastCommonKey] {
			leastCommonKey = k
		}
	}

	return leastCommonKey
}

func message(s [][]string, t string) string {
	var mostCommonMessage []string
	var leastCommonMessage []string

	transposed := transpose(s)

	for _, x := range transposed {
		mostCommonMessage = append(mostCommonMessage, mostCommon(x))
		leastCommonMessage = append(leastCommonMessage, leastCommon(x))
	}

	if t == "most" {
		return strings.Join(mostCommonMessage, "")
	} else {
		return strings.Join(leastCommonMessage, "")
	}
}

func parseInput(name string) [][]string {
	input := [][]string{}

	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		splitLine := strings.Split(line, "")

		input = append(input, splitLine)
	}

	return input
}

func main() {
	input := parseInput("input.txt")

	fmt.Printf("part 1: %s\n", message(input, "most"))
	fmt.Printf("part 2: %s\n", message(input, "least"))
}
