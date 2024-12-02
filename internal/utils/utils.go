package utils

import (
	"fmt"
	"math"
	"strings"
)

func InputAsLines(input []byte) []string {
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	return lines
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrintResults(part1, part2 any) {
	fmt.Printf("Part 1: %v\n", part1)
	fmt.Printf("Part 2: %v\n", part2)
}

func IntsOnLine(line string) []int {
	var ints []int
	for _, s := range strings.Fields(line) {
		var i int
		_, err := fmt.Sscanf(s, "%d", &i)
		PanicIfErr(err)
		ints = append(ints, i)
	}
	return ints
}

func AbsDiffInt(n1, n2 int) int {
	return int(math.Abs(float64(n1) - float64(n2)))
}

func Count[T any](vs []T, pred func(T) bool) int {
	count := 0
	for _, v := range vs {
		if pred(v) {
			count++
		}
	}
	return count
}
