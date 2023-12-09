package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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
	fmt.Printf("Par 1: %d\n", solve(rows))
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "reading input file")
	rows := parseInput(data)
	part1(rows)
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
