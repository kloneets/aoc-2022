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
	data := strings.Split(strings.ReplaceAll(ReadFile(fileName), "\r\n", "\n"), "\n")
	if data[len(data)-1] == "" { //remove empty line if exists
		data = data[:len(data)-1]
	}

	return data
}
