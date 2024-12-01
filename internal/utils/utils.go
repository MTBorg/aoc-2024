package utils

import (
	"fmt"
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
