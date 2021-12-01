package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func totalValid(s []string) int {
	count := 0
	for _, x := range s {
		if valid(x) {
			count++
		}
	}
	return count
}

func correctFormat(s string) bool {
	split := strings.Split(s, "")
	for i := 0; i < len(split)-3; i++ {
		if (split[i] == split[i+3]) && (split[i+1] == split[i+2]) && (split[i] != split[i+1]) {
			return true
		}
	}

	return false
}

func valid(s string) bool {
	re := regexp.MustCompile(`([a-z]+)`)
	matches := re.FindAllString(s, -1)

	inner, outer := false, false

	for i, x := range matches {
		if i%2 == 0 {
			outer = outer || correctFormat(x)
		} else {
			inner = inner || correctFormat(x)
		}
	}

	return outer && !inner
}

func parseInput(name string) []string {
	var input []string
	file, _ := os.Open(name)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}

func main() {
	input := parseInput("input.txt")

	count := totalValid(input)

	fmt.Printf("part 1: %d\n", count)
}
