package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/tw"
)

type Exercise func(io.Reader) (uint64, error)

var exercises = map[string]Exercise{
	"2025-12-01/1": aoc2025.Day1Part1,
	"2025-12-01/2": aoc2025.Day1Part2,

	"2025-12-02/1": aoc2025.Day2Part1,
	"2025-12-02/2": aoc2025.Day2Part2,

	"2025-12-03/1": aoc2025.Day3Part1,
	"2025-12-03/2": aoc2025.Day3Part2,

	"2025-12-04/1": aoc2025.Day4Part1,
	"2025-12-04/2": aoc2025.Day4Part2,

	"2025-12-05/1": aoc2025.Day5Part1,
	"2025-12-05/2": aoc2025.Day5Part2,

	"2025-12-06/1": aoc2025.Day6Part1,
	"2025-12-06/2": aoc2025.Day6Part2,

	"2025-12-07/1": aoc2025.Day7Part1,
	"2025-12-07/2": aoc2025.Day7Part2,
}

func main() {
	var d string
	var p uint

	flag.StringVar(&d, "date", "", "The date to run in ISO-8601 Date format (YYYY-mm-dd)")
	flag.UintVar(&p, "exercise", 0, "The exercise to run (must be 1 or 2)")

	flag.Parse()
	if p != 1 && p != 2 {
		log.Fatalf("Exercise #%d is not valid. Must be 1 or 2", p)
	}

	day, err := time.Parse("2006-01-02", d)
	if err != nil || day.Month() != time.December {
		log.Fatalf("Passed date is invalid or not in December")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current directory: %v", err)
	}

	inputFileName := filepath.Join(wd, fmt.Sprintf("aoc%d", day.Year()), "testdata", fmt.Sprintf("day%d.txt", day.Day()))

	in, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer in.Close()

	exercise, ok := exercises[fmt.Sprintf("%s/%d", d, p)]
	if !ok {
		log.Fatalf("Exercise %d for %s has not been implemented", p, d)
	}

	start := time.Now()
	result, err := exercise(in)
	end := time.Now()

	runtimeDuration := end.Sub(start)
	resultStr := fmt.Sprintf("%d", result)
	resultAlign := tw.AlignRight
	if err != nil {
		resultStr = fmt.Sprintf("exercise returned an error: %v", err)
		resultAlign = tw.AlignLeft
	}

	w := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRowAutoWrap(72),
		tablewriter.WithHeaderAlignment(tw.AlignCenter),
		tablewriter.WithRowAlignmentConfig(tw.CellAlignment{PerColumn: []tw.Align{
			tw.AlignCenter,
			tw.AlignCenter,
			resultAlign,
			tw.AlignCenter,
		},
		}))
	defer func() {
		w.Render()
	}()
	w.Header("Date", "Exercise", "Result", "Runtime")
	w.Append([]string{d, fmt.Sprintf("%d", p), resultStr, runtimeDuration.String()})
}
