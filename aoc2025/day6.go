package aoc2025

import (
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

func Day6Part1(in io.Reader) (uint64, error) {
	grid, operators, err := getGridAndOperators(in)
	if err != nil {
		return 0, err
	}

	var colTotals []uint64

	for col := range grid.Cols() {
		colTotal := operators[col]
		op := validOperators[colTotal]
		for row := range grid.Rows() {
			cellValue, err := grid.ValueAt(col, row)
			if err != nil {
				return 0, err
			}

			colTotal = op(colTotal, cellValue)
		}

		colTotals = append(colTotals, colTotal)
	}

	return slices.Reduce(colTotals, add, 0), nil
}

func Day6Part2(in io.Reader) (uint64, error) {
	rows, err := utils.InputAsLines(in, true)
	if err != nil {
		return 0, err
	}

	rows = slices.Filter(rows, func(s string) bool {
		return len(strings.TrimSpace(s)) > 0
	})

	maxX := len(rows[0]) - 1
	maxY := len(rows) - 1
	var colTotals []uint64
	var colOperands []uint64
	for x := maxX; x >= 0; x-- {
		operand := make([]byte, maxY)
		for y := range maxY {
			operand[y] = rows[y][x]
		}

		u, err := stringToUint(strings.TrimSpace(string(operand)))
		if err != nil {
			return 0, err
		}
		colOperands = append(colOperands, u)
		// if rows[-1][x] is an operator, calculate that total, THEN decrement x to skip the space which will occur next or
		// exit the loop if we're at x == 0. if it's a space, replacing the spaces with zeroes

		switch rows[maxY][x] {
		case '+':
			colTotals = append(colTotals, slices.Reduce(colOperands, add, 0))
		case '*':
			colTotals = append(colTotals, slices.Reduce(colOperands, mul, 1))
		case ' ':
			continue
		}

		colOperands = nil
		x--
	}

	return slices.Reduce(colTotals, add, 0), nil
}

type operator func(a, b uint64) uint64

func add(a, b uint64) uint64 {
	return a + b
}

func mul(a, b uint64) uint64 {
	return a * b
}

var validOperators = []operator{add, mul}

func getGridAndOperators(in io.Reader) (*utils.Matrix[uint64], []uint64, error) {
	rows, err := utils.InputAsLines(in, false)
	if err != nil {
		return nil, nil, err
	}

	operatorRow := rows[len(rows)-1]
	operators := slices.Map(strings.Fields(operatorRow), func(op string) uint64 {
		switch op {
		case "+":
			return 0
		case "*":
			return 1
		default:
			return math.MaxUint64
		}
	})

	matrix, err := utils.NewMatrix[uint64](rows[:len(rows)-1], stringToUint, nil)

	return matrix, operators, err
}

func stringToUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 0)
}
