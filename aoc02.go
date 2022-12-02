package main

import (
	"fmt"
	"strconv"

	"koko.lv/aoc"
)

func aoc02() {
	calorieMap := aoc.FileLines("inputs/011")

	totalCalories := 0
	totalMax := 0
	secondMax := 0
	thirdMax := 0

	for _, v := range calorieMap {
		if v == "" {
			if totalMax < totalCalories {
				secondMax = totalMax
				totalMax = totalCalories
			} else if secondMax < totalCalories {
				thirdMax = secondMax
				secondMax = totalCalories
			} else if thirdMax < totalCalories {
				thirdMax = totalCalories
			}
			totalCalories = 0
		} else {
			n, _ := strconv.Atoi(v)
			totalCalories += n
		}
	}
	if totalMax < totalCalories {
		secondMax = totalMax
		totalMax = totalCalories
	} else if secondMax < totalCalories {
		thirdMax = secondMax
		secondMax = totalCalories
	} else if thirdMax < totalCalories {
		thirdMax = totalCalories
	}
	fmt.Println("2. day. Three maxes:", totalMax+secondMax+thirdMax)

}
