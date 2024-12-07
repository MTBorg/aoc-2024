package main

import (
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/MTBorg/aoc-2024/internal/utils"
)

//go:embed input.txt
var input string

var rules = make(map[string]struct{})
var updates [][]int

func main() {
	parseInput()

	var unordered [][]int
	sum := 0
	for _, update := range updates {
		if isOrdered(update) {
			sum += update[len(update)/2]
		} else {
			unordered = append(unordered, update)
		}
	}
	part1 := sum

	sum = 0
	for _, update := range unordered {
		ordered := order(update)
		sum += ordered[len(update)/2]
	}
	part2 := sum

	utils.PrintResults(part1, part2)
}

func isOrdered(update []int) bool {
	for i, n := range update {
		for _, n2 := range update[i+1:] {
			if _, ok := rules[strconv.Itoa(n2)+"|"+strconv.Itoa(n)]; ok {
				return false
			}
		}
	}
	return true
}

func order(update []int) []int {
	lookup := map[int]int{}
	for i, n := range update {
		lookup[n] = 0
		for j, n2 := range update {
			if i == j {
				continue
			}
			_, ok := rules[strconv.Itoa(n)+"|"+strconv.Itoa(n2)]
			if ok {
				lookup[n] = lookup[n] + 1
			}
		}
	}

	var nums [][2]int
	for key, value := range lookup {
		nums = append(nums, [2]int{key, value})
	}

	slices.SortFunc(nums, func(a, b [2]int) int {
		if a[1] < b[1] {
			return 1
		} else {
			return -1
		}
	})

	var res []int
	for _, num := range nums {
		res = append(res, num[0])
	}
	return res
}

func parseInput() {
	lines := utils.InputAsLines([]byte(input))
	split := 0
	for i, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			split = i
			break
		}

		rules[line] = struct{}{}
	}

	for _, line := range lines[split+1:] {
		var update []int
		sn := strings.Split(line, ",")
		for _, s := range sn {
			n, err := strconv.Atoi(s)
			utils.PanicIfErr(err)
			update = append(update, n)
		}
		updates = append(updates, update)

	}
}
