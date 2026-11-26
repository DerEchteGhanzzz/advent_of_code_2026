import os

content = """
package %s

import (
	"strings"
)

func solveA(input []string) string {
	return strings.Join(input[:], "\\n")
}

func solveB(input []string) string {
	return strings.Join(input[:], "\\n")
}

func Solve() (string, func([]string) string, func([]string) string) {
	return "%s", func(input []string) string { return solveA(input) }, func(input []string) string { return solveB(input) }
}
"""

for i in range(1, 13):
    dir = f'day{i}'
    os.mkdir(dir)
    with open(f'{dir}/day{i}.go', 'w') as f:
        f.write(content.replace('%s', f'day{i}'))
    with open(f'{dir}/test-input.txt', 'w') as f:
        f.write('')
    with open(f'{dir}/input.txt', 'w') as f:
        f.write('')