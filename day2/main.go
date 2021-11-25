package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func partOne(file *os.File) {
	pad := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	x, y := 1, 1

	// pad[y][x]
	// pad[2][1] == 8

	fmt.Print("part 1: ")
	for scanner.Scan() {
		line := scanner.Text()

		instructions := strings.Split(line, "")

		for _, instruction := range instructions {
			switch instruction {
			case "U":
				if y > 0 {
					y = y - 1
				}
			case "D":
				if y < len(pad)-1 {
					y = y + 1
				}
			case "L":
				if x > 0 {
					x = x - 1
				}
			case "R":
				if x < len(pad[y])-1 {
					x = x + 1
				}
			}
		}

		fmt.Print(pad[y][x])
	}
	fmt.Println()
}

func partTwo(file *os.File) {
	pad := [][]rune{
		{' ', ' ', '1', ' ', ' '},
		{' ', '2', '3', '4', ' '},
		{'5', '6', '7', '8', '9'},
		{' ', 'A', 'B', 'C', ' '},
		{' ', ' ', 'D', ' ', ' '},
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	x, y := 0, 2

	fmt.Print("part 2: ")
	for scanner.Scan() {

		line := scanner.Text()

		instructions := strings.Split(line, "")

		for _, instruction := range instructions {
			switch instruction {
			case "U":
				if y > 0 && pad[y-1][x] != ' ' {
					y = y - 1
				}
			case "D":
				if y < len(pad)-1 && pad[y+1][x] != ' ' {
					y = y + 1
				}
			case "L":
				if x > 0 && pad[y][x-1] != ' ' {
					x = x - 1
				}
			case "R":
				if x < len(pad[y])-1 && pad[y][x+1] != ' ' {
					x = x + 1
				}
			}
		}

		fmt.Print(string(pad[y][x]))
	}
	fmt.Println()
}

func main() {
	file1, err := os.Open("input.txt")
	file2, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error")
	}

	defer file1.Close()
	defer file2.Close()

	partOne(file1)
	partTwo(file2)
}
