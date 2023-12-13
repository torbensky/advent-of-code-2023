package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
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

func (g grid) isColReflectionPoint(idx int) bool {
	for i, j := idx, idx+1; i >= 0 && j < g.width(); i, j = i-1, j+1 {
		if !slices.Equal(g.col(i), g.col(j)) {
			return false
		}
	}
	return true
}

func (g grid) isRowReflectionPoint(idx int) bool {
	for i, j := idx, idx+1; i >= 0 && j < g.height(); i, j = i-1, j+1 {
		if !slices.Equal(g.row(i), g.row(j)) {
			return false
		}
	}
	return true
}

func (g grid) findReflection() (row int, col int) {
	row, col = -1, -1
	for i := 0; i < g.width()-1; i++ {
		if g.isColReflectionPoint(i) {
			if col >= 0 {
				panic("multiple")
			}
			col = i
		}
	}
	for i := 0; i < g.height()-1; i++ {
		if g.isRowReflectionPoint(i) {
			if row >= 0 {
				panic("multiple")
			}
			row = i
		}
	}
	return row, col
}

func part1(patterns []grid) {
	answer := 0
	for _, p := range patterns {
		row, col := p.findReflection()
		if col >= 0 {
			answer += col + 1
		}
		if row >= 0 {
			answer += 100 * (row + 1)
		}
		// fmt.Println(p)
		fmt.Println(row, col, " => ", answer)
		// fmt.Scanln()
	}
	fmt.Printf("Part 1: %d\n", answer)
}

func main() {
	start := time.Now()
	patterns := parseInput()
	part1(patterns)
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
func must(err error, msg string) {
	if err != nil {
		slog.With("error", err).Error(msg)
	}
}
