
package day7

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
	return "7", func(input []string) string { return solveA(input) }, func(input []string) string { return solveB(input) }
}
