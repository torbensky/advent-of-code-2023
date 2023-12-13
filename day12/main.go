package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
)

const (
	operational byte = '.'
	damaged     byte = '#'
	unknown     byte = '?'
)

type record struct {
	conditions []byte
	checks     []int
}

func (r record) String() string {
	return fmt.Sprintf("%s %v", string(r.conditions), r.checks)
}

func parseRecord(line []byte) (result record) {
	// format is like "#..??.#?..?#.#.? 1,2,2"
	parts := bytes.Split(line, []byte{' '})
	result.conditions = make([]byte, len(parts[0]))
	copy(result.conditions, parts[0])

	vs := bytes.Split(parts[1], []byte{','})
	result.checks = make([]int, len(vs))
	for i, v := range vs {
		result.checks[i] = mustInt(v, fmt.Sprintf("expecting ints for record checks - %v is not int", v))
	}

	return result
}

func parseInput(data []byte) []record {
	lines := bytes.Split(data, []byte{'\n'})
	records := make([]record, len(lines))
	for i, line := range lines {
		records[i] = parseRecord(line)
	}
	return records
}

func (r record) status() (result []int) {
	dmgCount := 0
	for _, c := range r.conditions {
		if c == damaged {
			dmgCount++
			continue
		}
		if dmgCount > 0 {
			result = append(result, dmgCount)
			dmgCount = 0
		}
	}
	if dmgCount > 0 {
		result = append(result, dmgCount)
	}
	return result
}

func (r record) setCondition(pos int, c byte) record {
	conditions := make([]byte, len(r.conditions))
	copy(conditions, r.conditions)
	conditions[pos] = c
	return record{conditions, r.checks}
}

type stringset map[string]struct{}

func newStringSet(s string) stringset {
	return map[string]struct{}{s: {}}
}

func (ss stringset) addAll(others stringset) {
	for k, v := range others {
		ss[k] = v
	}
}

func findValidCombinations(r record, cache map[string]stringset) stringset {
	if results, ok := cache[r.String()]; ok {
		return results
	}
	if slices.Equal(r.checks, r.status()) {
		c := bytes.ReplaceAll(r.conditions, []byte{unknown}, []byte{operational})
		return newStringSet(string(c))
	}

	validSet := stringset{}
	for i, c := range r.conditions {
		if c == unknown {
			vc := findValidCombinations(r.setCondition(i, damaged), cache)
			validSet.addAll(vc)
		}
	}
	cache[r.String()] = validSet
	return validSet
}

func part1(records []record) {
	answer := 0
	for _, r := range records {
		answer += len(findValidCombinations(r, map[string]stringset{}))
	}
	fmt.Printf("Part 1: %d\n", answer)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading puzzle input")
	records := parseInput(data)
	part1(records)
}

func must(err error, msg string) {
	if err != nil {
		slog.With("error", err).Error(msg)
	}
}

func mustInt(b []byte, msg string) int {
	i, err := strconv.Atoi(string(b))
	must(err, msg)
	return i
}
