package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type span struct {
	start int
	end   int
}

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

func (m Mapping) applySpan(source span) (before *span, mapped *span, after *span) {
	start, end := m.src, m.src+m.len-1
	if source.start < start {
		before = &span{start: source.start, end: minInt(source.end, start-1)}
	}
	if source.end >= start && source.start <= end {
		s, _ := m.apply(maxInt(start, source.start))
		e, _ := m.apply(minInt(end, source.end))
		mapped = &span{start: s, end: e}
	}
	if source.end > end {
		after = &span{start: maxInt(end+1, source.start), end: source.end}
	}

	return before, mapped, after
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

func (ms Mappings) applySpan(source span) (result []span) {
	toMap := []span{source}
	for _, m := range ms {
		next := []span{}
		for _, s := range toMap {
			before, mapped, after := m.applySpan(s)
			if before != nil {
				next = append(next, *before)
			}
			if mapped != nil {
				result = append(result, *mapped)
			}
			if after != nil {
				next = append(next, *after)
			}
		}
		toMap = next
	}
	if len(toMap) > 0 {
		result = append(result, toMap...)
	}
	return result
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

func parseMappings(chunks []string) (result Almanac) {
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

type Almanac []Mappings

func (a Almanac) locations(source span) (result []span) {
	result = []span{source}
	for _, m := range a {
		next := []span{}
		for _, s := range result {
			next = append(next, m.applySpan(s)...)
		}
		result = next
	}
	return result
}

func part2(chunks []string) {
	seeds := parseSeeds(chunks[0])
	almanac := parseMappings(chunks[1:])
	lowest := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		locs := almanac.locations(span{start: seeds[i], end: seeds[i] + seeds[i+1] - 1})
		for _, s := range locs {
			if s.start < lowest {
				lowest = s.start
			}
		}
	}

	fmt.Printf("Part 2: %d\n", lowest)
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	must(err, "unable to read input file")
	chunks := strings.Split(strings.TrimSpace(string(data)), "\n\n")
	part1(chunks)
	part2(chunks)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
