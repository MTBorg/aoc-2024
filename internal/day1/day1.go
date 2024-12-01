package day1

import (
	"fmt"
	"math"
	"slices"

	"github.com/MTBorg/aoc-2024/internal/utils"
)

func Day1(input []byte) {
	lines := utils.InputAsLines(input)
	var left []int
	var right []int

	for _, line := range lines {
		var l, r int
		_, err := fmt.Sscanf(line, "%d%d", &l, &r)
		utils.PanicIfErr(err)
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	part1 := 0
	for i := range len(left) {
		part1 += int(math.Abs(float64(left[i] - right[i])))
	}

	part2 := 0
	for _, l := range left {
		count := 0
		for _, r := range right {
			if r == l {
				count++
			}
		}
		part2 += l * count
	}

	utils.PrintResults(part1, part2)
}
