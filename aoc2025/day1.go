package aoc2025

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

const (
	day1LockMinVal           = 0
	day1LockMaxVal           = 99
	day1LockRange            = day1LockMaxVal - day1LockMinVal + 1
	day1LockStartingPosition = day1LockMinVal + (day1LockRange / 2)
)

func Day1Part1(in io.Reader) (result int, err error) {
	moves, err := parseDay1Input(in)
	if err != nil {
		return
	}

	position := day1LockStartingPosition
	totalZeroes := 0
	for _, move := range moves {
		position += move + day1LockRange
		for position < 0 {
			position += day1LockRange
		}

		position %= day1LockRange

		if position == 0 {
			totalZeroes++
		}
	}

	result = totalZeroes
	return
}

func Day1Part2(in io.Reader) (result int, err error) {
	moves, err := parseDay1Input(in)
	if err != nil {
		return
	}

	position := day1LockStartingPosition
	totalZeroes := 0
	for _, move := range moves {
		var z int
		position, z = countZeroes(position, move)
		totalZeroes += z
	}

	result = totalZeroes
	return
}

func parseDay1Input(in io.Reader) (moves []int, err error) {
	var lines []string
	lines, err = utils.InputAsLines(in, false)
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	moves = slices.Map(lines, func(line string) int {
		normalized := strings.Map(func(r rune) rune {
			if r == 'L' {
				return '-'
			}

			if r == 'R' {
				return -1
			}

			return r
		}, strings.ToUpper(line))

		m, e := strconv.Atoi(normalized)
		if e != nil {
			panic(e)
		}

		return m
	})

	return
}

func countZeroes(startingPosition, move int) (endPosition, zeroes int) {
	counter := move / utils.Abs(move)

	endPosition = startingPosition

	for range utils.Abs(move) {
		endPosition += counter
		switch {
		case endPosition < day1LockMinVal:
			endPosition += day1LockRange
		case endPosition > day1LockMaxVal:
			endPosition = day1LockMinVal + (endPosition % day1LockRange)
		}

		if endPosition == day1LockMinVal {
			zeroes++
		}
	}

	return
}
