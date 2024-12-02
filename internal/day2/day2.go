package day2

import (
	"github.com/MTBorg/aoc-2024/internal/utils"
)

func Day2(inputData []byte) {
	lines := utils.InputAsLines(inputData)
	var rows [][]int

	for _, line := range lines {
		rows = append(rows, utils.IntsOnLine(line))
	}

	count := utils.Count(rows, isSafe)
	part1 := count

	count = utils.Count(rows, func(t []int) bool {
		return isSafe(t) || isSafeable(t)
	})
	part2 := count

	utils.PrintResults(part1, part2)
}

func isSafe(levels []int) bool {
	increasing := levels[0] < levels[1]
	sameDirection := func() bool {
		for i := range levels[:len(levels)-1] {
			if levels[i] < levels[i+1] != increasing {
				return false
			}
		}
		return true
	}

	validDelta := func() bool {
		for i := range levels[:len(levels)-1] {
			diff := utils.AbsDiffInt(levels[i], levels[i+1])
			if diff == 0 || diff > 3 {
				return false
			}
		}
		return true
	}

	res := sameDirection() && validDelta()
	return res
}

func isSafeable(report []int) bool {
	for i := range report {
		temp := make([]int, len(report)-1)
		copy(temp[:i], report[:i])
		copy(temp[i:], report[i+1:])
		if isSafe(temp) {
			return true
		}
	}
	return false
}
