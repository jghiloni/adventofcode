package aoc2025

import (
	"io"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

const day4RollChar = '@'

func Day4Part1(in io.Reader) (uint64, error) {
	grid, err := utils.NewByteGridFromInput(in)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	for y := range grid.Rows() {
		for x := range grid.Cols() {
			currentPosition := grid.ValueAt(x, y)
			val, valid := currentPosition.Value()
			if valid && val == day4RollChar {
				_, subtotal := grid.MatchingNeighbors(currentPosition, func(c utils.ByteGridCoordinate) bool {
					v, _ := c.Value()
					return v != day4RollChar
				})

				if subtotal > 4 {
					total++
				}
			}
		}
	}

	return total, nil
}

func Day4Part2(in io.Reader) (uint64, error) {
	panic("unimplemented")
}

func getSurroundingRollCounts(grid [][]byte) []int {
	counts := make([]int, 0, len(grid)*len(grid[0]))
	for y := range grid {
		for x := range grid[y] {
			counts = append(counts, getSurroundingRollCount(grid, y, x))
		}
	}

	return slices.Filter(counts, func(i int) bool {
		return i > 0
	})
}

func getSurroundingRollCount(grid [][]byte, row, col int) int {
	if grid[row][col] != day4RollChar {
		return 0
	}

	checks := [][]int{
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col - 1}, {row, col + 1},
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}

	count := 0
	for _, check := range checks {
		y, x := check[0], check[1]
		if 0 < y && y < len(grid) &&
			0 < x && x < len(grid[y]) {
			if grid[y][x] == day4RollChar {
				count++
			}
		}
	}

	return count
}
