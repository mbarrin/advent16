package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func expand(compressed string) int {
	fmt.Printf("compressed: %s\n", compressed)

	head, count, repeats, tail := "", 0, 0, ""
	end := regexp.MustCompile(`^\w+$`)

	if end.MatchString(compressed) {
		fmt.Printf("end of chain: %s\n", compressed)
		return len(compressed)
	}

	if strings.HasPrefix(compressed, "(") {
		fmt.Sscanf(compressed, "(%dx%d)%s", &count, &repeats, &tail)
		return repeats * expand(tail)
	}

	re := regexp.MustCompile(`(\w+)\((\d+)x(\d+)\)(.+)`)
	matches := re.FindAllStringSubmatch(compressed, -1)

	head = string(matches[0][1])
	count, _ = strconv.Atoi(string(matches[0][2]))
	repeats, _ = strconv.Atoi(string(matches[0][3]))
	tail = string(matches[0][4])

	return expand(head) + expand(tail)
}

func decryptV2(encrypted string) int {
	split := strings.Split(encrypted, "")

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

			end := endIndex + 1

			expander := strings.Join(split[startIndex:end], "")

			fmt.Sscanf(expander, "(%dx%d)", &count, &repeats)

			text := strings.Join(split[startIndex:end+count], "")

			outputLength += expand(text)
			i += len(expander) - 1 + count
		} else {
			outputLength++
		}
	}

	return outputLength
}

func decrypt(encrypted string) int {
	split := strings.Split(encrypted, "")

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

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Printf("part 1: %d\n", decrypt(scanner.Text()))
	}
}
