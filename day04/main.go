package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type card struct {
	num     int
	winning []string
	have    []string
}

var cardNumsRE = regexp.MustCompile(`\s+`)

func parseCard(line string) card {
	// Example: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	parts := strings.Split(line, ": ")
	num := num(cardNumsRE.Split(parts[0], -1)[1])
	parts = strings.Split(parts[1], " | ")
	winning := cardNumsRE.Split(strings.TrimSpace(parts[0]), -1)
	have := cardNumsRE.Split(strings.TrimSpace(parts[1]), -1)

	return card{num: num - 1, winning: winning, have: have}
}

func part1(lines []string) {
	result := 0
	for _, line := range lines {
		card := parseCard(line)
		numWon := len(intersect(card.winning, card.have))
		if numWon > 0 {
			result += int(math.Pow(2, float64(numWon-1)))
		}
	}
	fmt.Printf("Part 1: %d\n", result)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	part1(lines)
}

func intersect(s1, s2 []string) (result []string) {
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				result = append(result, v1)
				break
			}
		}
	}
	return result
}

func num(raw string) int {
	v, err := strconv.Atoi(raw)
	must(err)
	return v
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
