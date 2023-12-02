package main

import (
	"fmt"
	"log/slog"
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

var digNums = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func translate(digit string) string {
	switch digit {
	case "one", "1":
		return "1"
	case "two", "2":
		return "2"
	case "three", "3":
		return "3"
	case "four", "4":
		return "4"
	case "five", "5":
		return "5"
	case "six", "6":
		return "6"
	case "seven", "7":
		return "7"
	case "eight", "8":
		return "8"
	case "nine", "9":
		return "9"
	}
	slog.Error("Encountered invalid digit", "digit", digit)
	return digit // shouldn't happen
}

func findFirst(str string, substrs []string) string {
	firstIdx := len(str)
	value := ""
	for _, substr := range substrs {
		if pos := strings.Index(str, substr); pos > -1 && pos < firstIdx {
			firstIdx = pos
			value = substr
		}
	}
	return value
}
func findLast(str string, substrs []string) string {
	lastIdx := -1
	value := ""
	for _, substr := range substrs {
		if pos := strings.LastIndex(str, substr); pos > lastIdx {
			lastIdx = pos
			value = substr
		}
	}
	return value
}

func part2(lines []string) {
	sum := 0
	for _, line := range lines {
		first := translate(findFirst(line, digNums))
		last := translate(findLast(line, digNums))
		v, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	fmt.Println("part 2:", sum)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	part1(lines)
	part2(lines)
}
