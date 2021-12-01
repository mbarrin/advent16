package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func alreadyFilled(i []int, val int) bool {
	for _, x := range i {
		if x == val {
			return true
		}

	}
	return false
}

func main() {
	id := "abbhdwsy"
	//id := "abc"
	code := ""
	finalCode := ""
	secondCode := make([]byte, 8)
	checker := []int{}
	counter := 0

	for {
		counter += 1

		if len(code) == 8 {
			finalCode = code
		}

		if len(finalCode) == 8 && len(checker) == 8 {
			break
		}

		input := fmt.Sprintf("%s%d", id, counter)

		hash := md5.Sum([]byte(input))

		stringifiedHash := fmt.Sprintf("%x", hash)

		if strings.HasPrefix(stringifiedHash, "00000") {
			code = fmt.Sprintf("%s%s", code, string(stringifiedHash[5]))

			loc, err := strconv.Atoi(string(stringifiedHash[5]))

			if err != nil || loc >= 8 {
				continue
			}

			if alreadyFilled(checker, loc) {
				continue
			}

			checker = append(checker, loc)

			secondCode[loc] = stringifiedHash[6]
		}

	}

	fmt.Printf("part 1: %s\n", finalCode)
	fmt.Printf("part 2: %s\n", secondCode)
}
