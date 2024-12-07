package day4

import (
	"strings"

	"github.com/MTBorg/aoc-2024/internal/utils"
)

func Day4(inputData []byte) {
	lines := utils.InputAsLines(inputData)
	var chars [][]string
	for _, line := range lines {
		chars = append(chars, strings.Split(line, ""))
	}

	directions := directions()

	count := 0
	for y := range chars {
		for x := range chars[y] {
			for _, direction := range directions {
				if hasXmas(x, y, chars, direction) {
					count++
				}
			}
		}
	}
	part1 := count

	count = 0
	width := len(chars[0])
	height := len(chars)
	for y := range chars {
		for x := range chars[y] {
			if chars[y][x] != "A" {
				continue
			}
			if x < 1 || x > width-2 || y < 1 || y > height-2 {
				continue
			}

			case1 := (chars[y-1][x+1] == "S" && chars[y+1][x-1] == "M") ||
				(chars[y-1][x+1] == "M" && chars[y+1][x-1] == "S")
			case2 := (chars[y-1][x-1] == "S" && chars[y+1][x+1] == "M") ||
				(chars[y-1][x-1] == "M" && chars[y+1][x+1] == "S")
			if case1 && case2 {
				count++
			}
		}
	}
	part2 := count

	utils.PrintResults(part1, part2)
}

func hasXmas(x, y int, chars [][]string, direction [2]int) bool {
	width := len(chars[0])
	height := len(chars[1])
	dx, dy := direction[1], direction[0]

	for i, c := range "XMAS" {
		xi := x + i*dx
		yi := y + i*dy

		if xi < 0 || xi > width-1 || yi < 0 || yi > height-1 {
			return false
		}

		if chars[yi][xi] != string(c) {
			return false
		}
	}
	return true
}

func directions() [][2]int {
	return [][2]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // down
		{-1, 0},  // up
		{1, 1},   // down right
		{1, -1},  // down left
		{-1, 1},  // up right
		{-1, -1}, // up left
	}
}
