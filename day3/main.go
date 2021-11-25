package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func isValid(sides []string) bool {
	sidesInt := []int{}

	for _, x := range sides {
		tmp, _ := strconv.Atoi(x)
		sidesInt = append(sidesInt, tmp)
	}

	sort.Ints(sidesInt)

	if sidesInt[2] < sidesInt[0]+sidesInt[1] {
		return true
	}

	return false
}

func main() {
	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	re := regexp.MustCompile(`(\d+)`)

	count, otherCount := 0, 0
	col1 := []string{}
	col2 := []string{}
	col3 := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		split := re.FindAllString(line, -1)

		if isValid(split) {
			count++
		}

		col1 = append(col1, split[0])
		col2 = append(col2, split[1])
		col3 = append(col3, split[2])
	}

	for i := 0; i < len(col1); i = i + 3 {
		if isValid([]string{col1[i], col1[i+1], col1[i+2]}) {
			otherCount++
		}
		if isValid([]string{col2[i], col2[i+1], col2[i+2]}) {
			otherCount++
		}
		if isValid([]string{col3[i], col3[i+1], col3[i+2]}) {
			otherCount++
		}
	}

	fmt.Printf("part 1: %d\n", count)
	fmt.Printf("part 2: %d\n", otherCount)
}
