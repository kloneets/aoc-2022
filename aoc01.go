package main

import (
	"fmt"
	"strconv"

	"koko.lv/aoc"
)

func aoc01() {
	calorieMap := aoc.FileLines("inputs/011")

	totalCalories := 0
	totalMax := 0

	for _, v := range calorieMap {
		if v == "" {
			if totalMax < totalCalories {
				totalMax = totalCalories
			}
			totalCalories = 0
		} else {
			n, _ := strconv.Atoi(v)
			totalCalories += n
		}
	}
	if totalCalories > totalMax {
		totalMax = totalCalories
	}
	fmt.Println("1. day. Max calories:", totalMax)
}
