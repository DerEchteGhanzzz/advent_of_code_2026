package parsing

import (
	"os"
)

func ParseFile(day string, isTest ...bool) string {
	var fileName = "input.txt"
	if len(isTest) > 0 {
		fileName = "test-" + fileName
	}
	bytes, err := os.ReadFile(day + "/" + fileName)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
