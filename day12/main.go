package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type record struct {
	conditions string
	checks     []int
}

// turn consecutive .'s into one '.'
func normalizeDots(record string) string {
	var sb strings.Builder
	seen := false
	for _, c := range record {
		if c == '.' {
			if !seen {
				sb.WriteRune(c)
			}
			seen = true
			continue
		}
		seen = false
		sb.WriteRune(c)
	}
	return sb.String()
}

func (r record) optimize() record {
	lastNonOperational := strings.LastIndexAny(r.conditions, "?#")
	oc := normalizeDots(r.conditions[:lastNonOperational+1])
	return record{oc, r.checks}
}

func parseRecord(line []byte) record {
	// format is like "#..??.#?..?#.#.? 1,2,2"
	parts := bytes.Split(line, []byte{' '})

	vs := bytes.Split(parts[1], []byte{','})
	checks := make([]int, len(vs))
	for i, v := range vs {
		checks[i] = mustInt(v, fmt.Sprintf("expecting ints for record checks - %v is not int", v))
	}

	return record{string(parts[0]), checks}
}

func parseInput(data []byte) []record {
	lines := bytes.Split(data, []byte{'\n'})
	records := make([]record, len(lines))
	for i, line := range lines {
		records[i] = parseRecord(line)
	}
	return records
}

func (r record) repeat(n int) record {
	var sb strings.Builder
	var checks []int
	for i := 0; i < n; i++ {
		sb.WriteString(r.conditions)
		if i != n-1 {
			sb.WriteRune('?')
		}
		checks = append(checks, r.checks...)
	}
	return record{sb.String(), checks}
}

var cache = map[string]int{}

func countArrangements(record string, groups []int) int {
	key := fmt.Sprintf("%s%v", record, groups)
	if v, ok := cache[key]; ok {
		return v
	}

	// check if we consumed all groups and whether that is a match or not
	if len(groups) == 0 {
		if strings.ContainsRune(record, '#') {
			return 0
		}
		return 1
	}

	// ignore leading operational records
	start := 0
	for ; start < len(record)-1; start++ {
		if record[start] != '.' {
			break
		}
	}

	// consume the next group
	end := len(record) - len(groups) + 1 // space for '.' separators
	for _, g := range groups {
		end -= g // space for each damaged area
	}

	nextDmg := strings.IndexRune(record, '#')
	if nextDmg >= 0 {
		end = min(end, nextDmg)
	}
	count := 0
	for i := start; i < end+1; i++ {
		windowEnd := i + groups[0]
		window := record[i:windowEnd]
		if strings.Count(window, ".") > 0 {
			continue
		}

		next := record[windowEnd:]
		if strings.HasPrefix(next, "#") {
			continue
		}
		if len(groups) > 1 {
			next = next[1:]
		}
		count += countArrangements(next, groups[1:])
	}
	cache[key] = count
	return count
}

func part1(records []record) {
	answer := 0
	for _, r := range records {
		r = r.optimize()
		num := countArrangements(r.conditions, r.checks)
		answer += num
	}
	fmt.Printf("Part 1: %d\n", answer)
}

func part2(records []record) {
	answer := 0
	for _, r := range records {
		r = r.repeat(5).optimize()
		num := countArrangements(r.conditions, r.checks)
		answer += num
	}
	fmt.Printf("Part 2: %d\n", answer)
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading puzzle input")
	records := parseInput(data)
	part1(records)
	part2(records)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
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
