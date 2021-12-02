package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

func rotateRow(row []bool, amount int) []bool {
	for i := 0; i < amount; i++ {
		last := row[len(row)-1]
		row = row[:len(row)-1]
		row = append([]bool{last}, row...)
	}
	return row
}

func rotateColumn(grid [][]bool, x int, amount int) [][]bool {
	col := []bool{}
	for _, row := range grid {
		col = append(col, row[x])
	}

	col = rotateRow(col, amount)

	for i := range grid {
		grid[i][x] = col[i]
	}

	return grid
}

func turnOn(grid [][]bool, x, y int) [][]bool {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			grid[i][j] = true
		}
	}
	return grid
}

func code(grid [][]bool) {
	lookup := map[bool]string{true: "*", false: " "}

	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			fmt.Print(lookup[grid[y][x]])
			if x == 49 {
				fmt.Printf("\n")
			}
		}
	}
}

func lit(grid [][]bool) int {
	count := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				count++
			}

		}
	}

	return count
}

func main() {
	grid := make([][]bool, 6)

	for i := range grid {
		grid[i] = make([]bool, 50)
	}

	input := parseInput("input.txt")
	var y, x, amount int

	for _, command := range input {
		if strings.HasPrefix(command, "rect") {
			fmt.Sscanf(command, "rect %dx%d", &x, &y)
			grid = turnOn(grid, x, y)

		} else if strings.HasPrefix(command, "rotate row") {
			fmt.Sscanf(command, "rotate row y=%d by %d", &y, &amount)
			grid[y] = rotateRow(grid[y], amount)

		} else if strings.HasPrefix(command, "rotate column") {
			fmt.Sscanf(command, "rotate column x=%d by %d", &x, &amount)
			grid = rotateColumn(grid, x, amount)
		}
	}

	fmt.Printf("part 1: %d\n", lit(grid))
	code(grid)
}
