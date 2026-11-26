
package day3

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
	return "3", func(input []string) string { return solveA(input) }, func(input []string) string { return solveB(input) }
}
