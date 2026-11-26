package main

import (
	"aoc2026/day1"
	"aoc2026/day10"
	"aoc2026/day11"
	"aoc2026/day12"
	"aoc2026/day2"
	"aoc2026/day3"
	"aoc2026/day4"
	"aoc2026/day5"
	"aoc2026/day6"
	"aoc2026/day7"
	"aoc2026/day8"
	"aoc2026/day9"
	"aoc2026/parsing"
	"fmt"
	"strings"
	"time"
)

func parseAndGo(day string, solveA func([]string) string, solveB func([]string) string) {
	var input []string = strings.Split(parsing.ParseFile("day"+day), "\n")
	if len(input) == 1 && len(input[0]) == 0 {
		return
	}
	fmt.Printf("Day %s\n", day)
	var startA = time.Now()
	var a = solveA(input)
	var durationA = time.Since(startA)
	fmt.Printf("A: %s in %s\n", a, durationA)
	var startB = time.Now()
	var b = solveB(input)
	var durationB = time.Since(startB)
	fmt.Printf("B: %s in %s\n", b, durationB)
	solveB(input)
}

func main() {
	parseAndGo(day1.Solve())
	parseAndGo(day2.Solve())
	parseAndGo(day3.Solve())
	parseAndGo(day4.Solve())
	parseAndGo(day5.Solve())
	parseAndGo(day6.Solve())
	parseAndGo(day7.Solve())
	parseAndGo(day8.Solve())
	parseAndGo(day9.Solve())
	parseAndGo(day10.Solve())
	parseAndGo(day11.Solve())
	parseAndGo(day12.Solve())
}
