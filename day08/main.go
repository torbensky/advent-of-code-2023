package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
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

func navigate(node, target string, moves []byte, network Network) int {
	steps := 0
	for {
		m := moves[steps%len(moves)]
		if m == 'L' {
			node = network[node].left
		} else {
			node = network[node].right
		}
		steps++
		if strings.HasSuffix(node, target) {
			return steps
		}
	}
}

func part1(moves []byte, network Network) {
	fmt.Printf("Part 1: %d\n", navigate("AAA", "ZZZ", moves, network))
}

func part2(moves []byte, network Network) {
	var steps []int
	for n := range network {
		if strings.HasSuffix(n, "A") {
			steps = append(steps, navigate(n, "Z", moves, network))
		}
	}
	fmt.Printf("Part 2: %d\n", lcmAll(steps[0], steps[1:]...))
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading input file")
	m, n := parseInput(bytes.Split(data, []byte{'\n'}))
	part1(m, n)
	part2(m, n)
}

func must(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %e", msg, err)
	}
}

// gcd,lcm,lcmAll inspired by:
// https://stackoverflow.com/questions/31302054/how-to-find-the-least-common-multiple-of-a-range-of-numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
func lcmAll(a int, bs ...int) int {
	result := a
	for _, b := range bs {
		result = lcm(result, b)
	}

	return result
}
