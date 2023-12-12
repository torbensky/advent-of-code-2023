package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
)

func emptyCols(grid [][]byte) (result []int) {
	for col := 0; col < len(grid[0]); col++ {
		empty := true
		for row := range grid {
			if grid[row][col] == '#' {
				empty = false
				break
			}
		}
		if empty {
			result = append(result, col)
		}
	}
	return result
}

func emptyRows(grid [][]byte) (result []int) {
	for i, row := range grid {
		if bytes.Count(row, []byte{'#'}) == 0 {
			result = append(result, i)
		}
	}
	return result
}

type point struct {
	x int
	y int
}

func findGalaxies(grid [][]byte) (galaxies []point) {
	for y, row := range grid {
		for x, tile := range row {
			if tile == '#' {
				galaxies = append(galaxies, point{x, y})
			}
		}
	}
	return galaxies
}

func solve(galaxies []point, emptyRows, emptyCols []int, expansionFactor int) int {
	gcopy := make([]point, len(galaxies))
	copy(gcopy, galaxies)
	for i, row := range emptyRows {
		for gi := range gcopy {
			if gcopy[gi].y > row+i*expansionFactor {
				gcopy[gi].y += expansionFactor
			}
		}
	}
	for i, col := range emptyCols {
		for gi := range gcopy {
			if gcopy[gi].x > col+i*expansionFactor {
				gcopy[gi].x += expansionFactor
			}
		}
	}
	answer := 0
	for i, g1 := range gcopy {
		for _, g2 := range gcopy[i+1:] {
			answer += abs(g1.x-g2.x) + abs(g1.y-g2.y)
		}
	}
	return answer
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading puzzle input")
	grid := bytes.Split(data, []byte{'\n'})
	emptyRows := emptyRows(grid)
	emptyCols := emptyCols(grid)
	galaxies := findGalaxies(grid)
	fmt.Printf("Part 1: %d\n", solve(galaxies, emptyRows, emptyCols, 1))
}

func must(err error, msg string) {
	if err != nil {
		slog.With("error", err).Error(msg)
	}
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
