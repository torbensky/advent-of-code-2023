package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseSeeds(line string) (seeds []int) {
	for _, s := range strings.Fields(line)[1:] {
		v, err := strconv.Atoi(s)
		must(err, "seed was not an integer")
		seeds = append(seeds, v)
	}
	return seeds
}

type Mapping struct {
	src int
	dst int
	len int
}

func (m Mapping) apply(source int) (int, bool) {
	diff := source - m.src
	if diff >= 0 && diff < m.len {
		return m.dst + diff, true
	}
	return -1, false
}

type Mappings []Mapping

func (ms Mappings) apply(source int) int {
	for _, m := range ms {
		if val, ok := m.apply(source); ok {
			return val
		}
	}
	return source
}

func parseMapping(line string) Mapping {
	parts := strings.Fields(line)
	dst, err := strconv.Atoi(parts[0])
	must(err, "destination not a valid int")
	src, err := strconv.Atoi(parts[1])
	must(err, "source not a valid int")
	length, err := strconv.Atoi(parts[2])
	must(err, "length not a valid int")
	return Mapping{dst: dst, src: src, len: length}
}

func parseMappings(chunks []string) (result []Mappings) {
	for _, chunk := range chunks {
		var m Mappings
		for _, line := range strings.Split(chunk, "\n")[1:] {
			m = append(m, parseMapping(line))
		}
		result = append(result, m)
		m = nil
	}
	return result
}

func part1(chunks []string) {
	seeds := parseSeeds(chunks[0])
	mappings := parseMappings(chunks[1:])
	lowest := math.MaxInt
	for _, seed := range seeds {
		val := seed
		for _, m := range mappings {
			val = m.apply(val)
		}
		if val < lowest {
			lowest = val
		}
	}
	fmt.Printf("Part 1: %d\n", lowest)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "unable to read input file")
	chunks := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	part1(chunks)
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
