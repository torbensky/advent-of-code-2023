package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func winsPossible(time int, record int) int {
	// distance function
	d := func(holdTime int) int {
		return holdTime * (time - holdTime)
	}

	// In order to avoid checking all possible hold times, we're going to solve
	// the quadratic equation that gives the distance traveled for any hold time
	//
	// i.e. for distance (d), hold time (x), total time (t) and record time (r):
	// r > d (which can be r = d + 1 since we only need to beat the record by 1)
	// r = x * (t - x) + 1
	// r = -x^2 + tx + 1
	// 0 = -x^2 + tx - r + 1

	// to solve the quadratic equation (ax^2 + bx + c), first we find the discriminant D:
	// a = -1, b = t, c = (-r + 1)
	// D = b^2 - 4ac
	D := time*time - 4*-1*(-1*record+1)

	// now we can solve for x using the formula
	// x = (-b +/- sqrt(D)) / 2a
	sD := math.Sqrt(float64(D))
	shortest := math.Round((float64(-time) + sD) / -2)
	longest := math.Round((float64(-time) - sD) / -2)

	// and because we rounded we may be off by a bit
	// back off by 1 for longest and shortest to be on safe side
	shortest += 1
	longest -= 1
	for {
		t := shortest - 1
		if d(int(t)) > record {
			shortest = t
		} else {
			break
		}
	}

	for {
		t := longest + 1
		if d(int(t)) > record {
			longest = t
		} else {
			break
		}
	}

	return int(longest - shortest + 1)
}

func part1(lines []string) {
	ts := strings.Fields(lines[0])[1:]
	ds := strings.Fields(lines[1])[1:]
	result := 1
	for i := range ts {
		t := mustInt(ts[i], "expecting int for a time value")
		d := mustInt(ds[i], "expecting int for a distance value")
		result *= winsPossible(t, d)
	}
	fmt.Printf("Part 1: %d\n", result)
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	must(err, "unable to read input file")
	lines := strings.Split(string(data), "\n")
	part1(lines)
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

func mustInt(str, msg string) int {
	i, err := strconv.Atoi(str)
	must(err, msg)
	return i
}

func must(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}
