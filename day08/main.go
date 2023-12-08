package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

type Network map[string]Node
type Node struct {
	left  string
	right string
}

func parseInput(lines [][]byte) (moves []byte, network Network) {
	moves = lines[0]
	network = Network{}
	for _, line := range lines[2:] {
		network[string(line[0:3])] = Node{string(line[7:10]), string(line[12:15])}
	}
	return moves, network
}

func part1(moves []byte, network Network) {
	steps := 0
	node := "AAA"
	target := "ZZZ"
	for {
		m := moves[steps%len(moves)]
		if m == 'L' {
			node = network[node].left
		} else {
			node = network[node].right
		}
		steps++
		if node == target {
			break
		}
	}
	fmt.Printf("Part 1: %d\n", steps)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading input file")
	m, n := parseInput(bytes.Split(data, []byte{'\n'}))
	part1(m, n)
}

func must(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %e", msg, err)
	}
}
