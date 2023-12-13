package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"
)

type grid [][]bool

func (g grid) String() string {
	var sb strings.Builder
	for i := 0; i < g.height(); i++ {
		for j := 0; j < g.width(); j++ {
			if g[i][j] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
func (g grid) row(n int) []bool {
	return g[n]
}
func (g grid) col(n int) []bool {
	col := make([]bool, g.height())
	for i := 0; i < g.height(); i++ {
		col[i] = g[i][n]
	}
	return col
}
func (g grid) width() int {
	return len(g[0])
}
func (g grid) height() int {
	return len(g)
}

func parseRow(row string) []bool {
	result := make([]bool, len(row))
	for i, c := range row {
		if c == '#' {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func parseInput() (result []grid) {
	file, err := os.Open(os.Args[1])
	must(err, "reading puzzle input")
	scanner := bufio.NewScanner(file)
	g := grid{}
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			result = append(result, g)
			g = grid{}
			continue
		}
		g = append(g, parseRow(row))
	}
	result = append(result, g)
	must(scanner.Err(), "problem scanning input file")
	return result
}

func diff(a, b []bool) (diffs int) {
	// NOTE: assuming a and b are the same length
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs++
		}
	}
	return diffs
}

func (g grid) isColReflectionPoint(idx int, smudge bool) bool {
	smudgeAmount := 0
	if smudge {
		smudgeAmount = 1
	}
	numDifferences := 0
	for i, j := idx, idx+1; i >= 0 && j < g.width(); i, j = i-1, j+1 {
		numDifferences += diff(g.col(i), g.col(j))
		if numDifferences > smudgeAmount {
			return false
		}
	}
	return numDifferences == smudgeAmount
}

func (g grid) isRowReflectionPoint(idx int, smudge bool) bool {
	smudgeAmount := 0
	if smudge {
		smudgeAmount = 1
	}
	numDifferences := 0
	for i, j := idx, idx+1; i >= 0 && j < g.height(); i, j = i-1, j+1 {
		numDifferences += diff(g.row(i), g.row(j))
		if numDifferences > smudgeAmount {
			return false
		}
	}
	return numDifferences == smudgeAmount
}

func (g grid) findReflection(smudge bool) (row int, col int) {
	for i := 0; i < g.width()-1; i++ {
		if g.isColReflectionPoint(i, smudge) {
			return -1, i
		}
	}
	for i := 0; i < g.height()-1; i++ {
		if g.isRowReflectionPoint(i, smudge) {
			return i, -1
		}
	}
	panic("No reflection found when one was expected")
}

func part1(patterns []grid) {
	answer := 0
	for _, p := range patterns {
		row, col := p.findReflection(false)
		if col >= 0 {
			answer += col + 1
		}
		if row >= 0 {
			answer += 100 * (row + 1)
		}
	}
	fmt.Printf("Part 1: %d\n", answer)
}

func part2(patterns []grid) {
	answer := 0
	for _, p := range patterns {
		row, col := p.findReflection(true)
		if col >= 0 {
			answer += col + 1
		}
		if row >= 0 {
			answer += 100 * (row + 1)
		}
	}
	fmt.Printf("Part 2: %d\n", answer)
}

func main() {
	start := time.Now()
	patterns := parseInput()
	part1(patterns)
	part2(patterns)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
func must(err error, msg string) {
	if err != nil {
		slog.With("error", err).Error(msg)
	}
}
