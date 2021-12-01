package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func totalValid(s []string) (int, int) {
	countTLS, countSSL := 0, 0
	for _, x := range s {
		if validTLS(x) {
			countTLS++
		}
		if validSSL(x) {
			countSSL++
		}
	}
	return countTLS, countSSL
}

func correctTLSFormat(s string) bool {
	split := strings.Split(s, "")
	for i := 0; i < len(split)-3; i++ {
		if (split[i] == split[i+3]) && (split[i+1] == split[i+2]) && (split[i] != split[i+1]) {
			return true
		}
	}

	return false
}

func validTLS(s string) bool {
	re := regexp.MustCompile(`([a-z]+)`)
	matches := re.FindAllString(s, -1)

	inner, outer := false, false

	for i, x := range matches {
		if i%2 == 0 {
			outer = outer || correctTLSFormat(x)
		} else {
			inner = inner || correctTLSFormat(x)
		}
	}

	return outer && !inner
}

func correctSSLFormats(s string) []string {
	aba := []string{}
	split := strings.Split(s, "")
	for i := 0; i < len(split)-2; i++ {
		if (split[i] == split[i+2]) && (split[i] != split[i+1]) {
			aba = append(aba, strings.Join(split[i:i+3], ""))
		}
	}
	return aba
}

func validSSL(s string) bool {
	re := regexp.MustCompile(`([a-z]+)`)
	matches := re.FindAllString(s, -1)

	abas := make(map[string]int)

	for i := 0; i <= len(matches); i += 2 {
		for _, x := range correctSSLFormats(matches[i]) {
			_, contains := abas[x]
			if !contains {
				abas[x]++
			}
		}

	}

	for i := 1; i < len(matches); i += 2 {
		for _, x := range correctSSLFormats(matches[i]) {
			_, contains := abas[swap(x)]
			if contains {
				return true
			}
		}

	}

	return false
}

func swap(s string) string {
	tmp := strings.Split(s, "")

	foo := []string{tmp[1], tmp[0], tmp[1]}

	return strings.Join(foo, "")
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

	tls, ssl := totalValid(input)

	fmt.Printf("part 1: %d\n", tls)
	fmt.Printf("part 2: %d\n", ssl)
}
