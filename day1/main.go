package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x float64
	y float64
}

func contains(s []coord, c coord) bool {
	for _, a := range s {
		if a.x == c.x && a.y == c.y {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error")
	}

	defer file.Close()

	tracker := coord{x: 0, y: 0}

	var (
		visited []coord
		seen    []coord
		tmp     coord
	)

	re := regexp.MustCompile("([A-Za-z])([0-9]+),?")

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	facing := "n"

	for scanner.Scan() {
		info := scanner.Text()

		split := re.FindAllStringSubmatch(info, -1)

		direction := split[0][1]
		count, _ := strconv.ParseFloat(split[0][2], 32)

		switch facing {
		case "n":
			if direction == "L" {
				for i := tracker.x; i > tracker.x-count; i-- {
					tmp = coord{i, tracker.y}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.x = tracker.x - count
				facing = "w"
			} else {
				for i := tracker.x; i < tracker.x+count; i++ {
					tmp = coord{i, tracker.y}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.x = tracker.x + count
				facing = "e"
			}
		case "e":
			if direction == "L" {
				for i := tracker.y; i < tracker.y+count; i++ {
					tmp = coord{tracker.x, i}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.y = tracker.y + count
				facing = "n"
			} else {
				for i := tracker.y; i > tracker.y-count; i-- {
					tmp = coord{tracker.x, i}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.y = tracker.y - count
				facing = "s"
			}
		case "s":
			if direction == "L" {
				for i := tracker.x; i < tracker.x+count; i++ {
					tmp = coord{i, tracker.y}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.x = tracker.x + count
				facing = "e"
			} else {
				for i := tracker.x; i > tracker.x-count; i-- {
					tmp = coord{i, tracker.y}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.x = tracker.x - count
				facing = "w"
			}
		case "w":
			if direction == "L" {
				for i := tracker.y; i > tracker.y-count; i-- {
					tmp = coord{tracker.x, i}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.y = tracker.y - count
				facing = "s"
			} else {
				for i := tracker.y; i < tracker.y+count; i++ {
					tmp = coord{tracker.x, i}
					if contains(visited, tmp) {
						seen = append(seen, tmp)
					}
					visited = append(visited, tmp)
				}

				tracker.y = tracker.y + count
				facing = "n"
			}
		}

	}

	fmt.Printf("part 1: %v\n", math.Abs(tracker.x)+math.Abs(tracker.y))
	fmt.Printf("part 2: %v\n", math.Abs(seen[0].x)+math.Abs(seen[0].y))
}
