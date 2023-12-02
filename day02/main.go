package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var gameRE = regexp.MustCompile(`Game (\d+): (.+)`)

type round struct {
	r int
	g int
	b int
}

func (r round) valid(maxR, maxG, maxB int) bool {
	return r.r <= maxR && r.g <= maxG && r.b <= maxB
}

type game struct {
	ID     int
	rounds []round
}

func parseRound(rstr string) round {
	round := round{}
	for _, r := range strings.Split(strings.TrimSpace(rstr), ",") {
		r = strings.TrimSpace(r)
		pair := strings.Split(r, " ")
		val, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		switch pair[1] {
		case "red":
			round.r = val
		case "green":
			round.g = val
		case "blue":
			round.b = val
		}
	}

	return round
}

func parseRounds(rstr string) (rounds []round) {
	rs := strings.Split(rstr, ";")
	for _, r := range rs {
		rounds = append(rounds, parseRound(r))
	}
	return rounds
}

func parseGame(line string) game {
	matches := gameRE.FindStringSubmatch(line)
	ID, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	rounds := parseRounds(matches[2])
	return game{ID: ID, rounds: rounds}
}

func part1(lines []string) {
	maxR, maxG, maxB := 12, 13, 14
	result := 0
GAMELOOP:
	for _, line := range lines {
		g := parseGame(line)
		for _, r := range g.rounds {
			if !r.valid(maxR, maxG, maxB) {
				continue GAMELOOP
			}
		}
		result += g.ID
	}
	fmt.Println("Part 1:", result)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	part1(lines)
}
