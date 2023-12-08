package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

type round struct {
	hand     []byte
	bid      int
	handType int
}

type cardCounter map[byte]int

func countCards(cards []byte) cardCounter {
	counts := map[byte]int{}
	for _, c := range cards {
		counts[c]++
	}
	return counts
}
func (cc cardCounter) getMostCommon() (b byte, c int) {
	for k, v := range cc {
		if v > c {
			b = k
			c = v
		}
	}
	return b, c
}

func handType(hand []byte) int {
	counts := countCards(hand)
	numGroups := len(counts)
	b, mostCommon := counts.getMostCommon()
	delete(counts, b)
	_, nextMostCommon := counts.getMostCommon()
	switch numGroups {
	case 1:
		// five of a kind (highest)
		return 7
	case 2:
		if mostCommon == 4 {
			// four of a kind
			return 6
		}
		// full house
		return 5
	case 3:
		if mostCommon == 3 && nextMostCommon == 1 {
			// three of a kind
			return 4
		}
		// two pair
		return 3
	case 4:
		// one pair
		return 2
	default:
		return 1
	}
}

func roundSortFunc(a, b round) int {
	// first sort by hand type
	if a.handType != b.handType {
		return a.handType - b.handType
	}

	// then sort by the card values, left->right (we can rely on ASCII compare)
	return bytes.Compare(a.hand, b.hand)
}

func hToR(hand, bid []byte) round {
	// we're going to convert the card faces to hex values according to their rank
	// this will give us the benefit of a naturally sortable / comparable hand
	for i, c := range hand {
		switch c {
		case 'T':
			hand[i] = 'A'
		case 'J':
			hand[i] = 'B'
		case 'Q':
			hand[i] = 'C'
		case 'K':
			hand[i] = 'D'
		case 'A':
			hand[i] = 'E'
		default:
			// leave as-is
		}
	}
	return round{hand, mustInt(bid, "bid should be an int"), handType(hand)}
}

func parseInput(lines [][]byte) []round {
	rounds := make([]round, len(lines))
	for i, line := range lines {
		parts := bytes.Split(line, []byte{' '})
		rounds[i] = hToR(parts[0], parts[1])
	}
	return rounds
}

func part1(rounds []round) {
	slices.SortFunc(rounds, roundSortFunc)
	answer := 0
	for i, r := range rounds {
		answer += (i + 1) * r.bid
	}
	fmt.Printf("Part 1: %d\n", answer)
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	must(err, "unable to read input file")
	rounds := parseInput(bytes.Split(data, []byte{'\n'}))
	part1(rounds)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func mustInt(b []byte, msg string) int {
	i, err := strconv.Atoi(string(b))
	must(err, msg)
	return i
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
