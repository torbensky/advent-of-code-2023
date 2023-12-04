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

func (c card) won() []string {
	return intersect(c.winning, c.have)
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

func part2(lines []string) {
	cards := make([]card, len(lines))
	for i, line := range lines {
		cards[i] = parseCard(line)
	}

	// cache the set of cards that are won by each card
	winCache := map[int][]card{}
	for _, c := range cards {
		numWon := len(c.won())
		for i := 1; i <= numWon; i++ {
			winCache[c.num] = append(winCache[c.num], cards[c.num+i])
		}
	}

	// going to use some dynamic programming to memoize how many cards are won
	// so we can instantly return the result when we repeat cards
	counterCache := map[int]int{}
	var counter func(c card) int
	counter = func(c card) int {
		// see if we've already counted all the cards that are won for this card
		if won, ok := counterCache[c.num]; ok {
			return won
		}

		// haven't counted the total cards won for this card yet

		// look up all the cards we win from the current card
		wonCards := winCache[c.num]

		// and add up the cards won from each of those to get the total for this card
		total := 0
		for _, cw := range wonCards {
			count := counter(cw)
			counterCache[cw.num] = count
			total += count + 1 // add one to include current card
		}
		counterCache[c.num] = total
		return total
	}

	result := 0
	for _, c := range cards {
		result += counter(c) + 1 // add one to include current card
	}
	fmt.Printf("Part 2: %d\n", result)
}

func part1(lines []string) {
	result := 0
	for _, line := range lines {
		card := parseCard(line)
		numWon := len(card.won())
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
	part2(lines)
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
