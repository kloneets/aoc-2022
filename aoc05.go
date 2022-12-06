package main

import (
	"fmt"

	"koko.lv/aoc"
)

type Step = int

const (
	Positions Step = iota
	Instruction
)

func aoc05() {
	data := aoc.FileLines("inputs/050")

	fmt.Println("Day 5, part 1. Message:", aoc05Part1(data))
	fmt.Println("Day 5, part 2. Message:", aoc05Part2(data))
}

func aoc05Part1(data []string) string {
	message := ""
	curStep := Positions
	var curPositions []string
	var instructions []string
	for _, v := range data {
		switch v {
		case "":
			curStep = Instruction
		default:
			if curStep == Positions {
				curPositions = append(curPositions, v)
			} else {
				instructions = append(instructions, v)
			}
		}
	}
	return message
}

func makeMatrix(positions []string) map[int]int {
	positions = positions[:len(positions)-1]
	matrix := make(map[int]int)
	// maxElements := (len(positions[0]) + 1) / 3
	// for _, v := range positions {

	// }
	return matrix
}

func aoc05Part2(data []string) string {
	return "todo"
}
