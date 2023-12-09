package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func parseInput(data []byte) (rows [][]int) {
	lines := bytes.Split(data, []byte{'\n'})
	rows = make([][]int, len(lines))
	for i := range lines {
		nums := bytes.Fields(lines[i])
		rows[i] = make([]int, len(nums))
		for j, n := range nums {
			rows[i][j] = mustInt(n, "rows should contain valid numbers only")
		}
	}
	return rows
}

func nextRow(row []int) (result []int, zeros bool) {
	if len(row) <= 1 {
		return []int{0}, true
	}

	result = make([]int, len(row)-1)
	prev := row[0]
	zeros = prev == 0
	for i, next := range row[1:] {
		result[i] = next - prev
		zeros = zeros && result[i] == 0
		prev = next
	}
	return result, zeros
}

func solveRow(row []int) (result int) {
	zeroes := false
	result = row[len(row)-1]
	for !zeroes {
		row, zeroes = nextRow(row)
		result += row[len(row)-1]
	}
	return result
}

func solve(rows [][]int) (result int) {
	for i := range rows {
		result += solveRow(rows[i])
	}
	return result
}

func part1(rows [][]int) {
	fmt.Printf("Part 1: %d\n", solve(rows))
}

func part2(rows [][]int) {
	for _, row := range rows {
		slices.Reverse(row)
	}
	fmt.Printf("Part 2: %d\n", solve(rows))
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading input file")
	rows := parseInput(data)
	part1(rows)
	// We can take advantage of an interesting property of the number pyramid to
	// solve part 2 and simply reverse the rows to get the answer.
	//
	// In part 1:
	// For a row of numbers v1, v2... vn, the last number of the next row is
	// given by vn - v(n-1) and then the next number for that row is v(n+1) = (vn - v(n-1) + vn).
	//
	// In part 2:
	// For the same row of numbers, the first number of the next row is given by
	// (v2 - v1) and then the previous number for that row is v0 = v1 - (v2 - v1)
	// When we reverse, v0 = v(n+1), v1 = vn, v2 = v(n-1) which means
	// v0 = v1 - (v2 - v1)
	// = v(n+1) = vn - (v(n-1) - vn)
	// = vn - v(n-1) + vn <--------- gee, doesn't that look familiar!? ;)
	part2(rows)
}

func mustInt(b []byte, msg string) int {
	i, err := strconv.Atoi(string(b))
	must(err, msg)
	return i
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
