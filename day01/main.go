package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitRegex = regexp.MustCompile(`(\d)`)

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		digits := digitRegex.FindAllString(line, -1)
		calibrationVal := digits[0] + digits[len(digits)-1]
		v, err := strconv.Atoi(calibrationVal)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Println("Part 1:", sum)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	part1(strings.Split(strings.TrimSpace(string(data)), "\n"))
}
