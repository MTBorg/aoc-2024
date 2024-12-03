package day3

import (
	"fmt"

	"github.com/MTBorg/aoc-2024/internal/utils"
)

func Day3(inputData []byte) {
	s := string(inputData)
	part1 := part1(s)
	part2 := part2(s)
	utils.PrintResults(part1, part2)
}

func part1(s string) int {
	res := 0
	for len(s) > 2 {
		if s[:3] == "mul" {
			var a, b int
			_, err := fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			if err == nil {
				res += a * b
			}
		}
		s = s[1:]
	}

	return res
}

func part2(s string) int {
	res := 0
	enabled := true

	for len(s) > 0 {
		if len(s) >= 7 && s[:7] == "don't()" {
			enabled = false
		} else if len(s) >= 4 && s[:4] == "do()" {
			enabled = true
		} else if len(s) >= 3 && s[:3] == "mul" && enabled {
			var a, b int
			_, err := fmt.Sscanf(s, "mul(%d,%d)", &a, &b)
			if err == nil {
				res += a * b
			}
		}
		s = s[1:]
	}

	return res
}
