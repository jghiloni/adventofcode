package aoc2025

import (
	"io"

	"github.com/jghiloni/adventofcode/utils"
)

const day4RollChar = '@'

func Day4Part1(in io.Reader) (uint64, error) {
	grid, err := utils.NewByteGridFromInput(in)
	if err != nil {
		return 0, err
	}

	total := getPassTotal(grid)

	return total, nil
}

func Day4Part2(in io.Reader) (uint64, error) {
	grid, err := utils.NewByteGridFromInput(in)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	subtotal := uint64(1)

	for i := 1; subtotal > 0; i++ {
		//fmt.Printf("Room before pass %d:\n%s\n\n", i, grid)
		subtotal = getPassTotal(grid)
		total += subtotal
	}

	return total, nil
}

func getPassTotal(grid *utils.ByteGrid) uint64 {
	total := uint64(0)
	removed := []utils.ByteGridCoordinate{}
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
					removed = append(removed, currentPosition)
				}
			}
		}
	}

	for _, c := range removed {
		grid.SetValueAt(c, 'x')
	}

	return total
}
