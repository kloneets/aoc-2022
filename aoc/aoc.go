package aoc

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	return str
}

func FileLines(fileName string) []string {
	return strings.Split(strings.ReplaceAll(ReadFile(fileName), "\r\n", "\n"), "\n")
}
