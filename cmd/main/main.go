package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MTBorg/aoc-2024/internal/day1"
	"github.com/MTBorg/aoc-2024/internal/day2"
	"github.com/MTBorg/aoc-2024/internal/day3"
	"github.com/MTBorg/aoc-2024/internal/day4"
)

type DayFunc = func(inputData []byte)

var days = map[string]DayFunc{
	"1": day1.Day1,
	"2": day2.Day2,
	"3": day3.Day3,
	"4": day4.Day4,
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("missing day")
	}
	day := os.Args[1]

	f, err := os.Open(fmt.Sprintf("internal/day%s/input.txt", day))
	if err != nil {
		panic(err)
	}

	inputData, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	days[day](inputData)
}
