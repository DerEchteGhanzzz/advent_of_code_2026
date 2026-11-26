
package day10

import (
	"strings"
)

func solveA(input []string) string {
	return strings.Join(input[:], "\n")
}

func solveB(input []string) string {
	return strings.Join(input[:], "\n")
}

func Solve() (string, func([]string) string, func([]string) string) {
	return "day10", func(input []string) string { return solveA(input) }, func(input []string) string { return solveB(input) }
}
