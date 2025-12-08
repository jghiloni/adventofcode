package aoc2025

import (
	"fmt"
	"io"
	"maps"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

func Day7Part1(in io.Reader) (uint64, error) {
	g, err := utils.NewByteGridFromInput(in)
	if err != nil {
		return 0, err
	}

	beamXs := make(map[int]bool)
	splitters := make(map[int]utils.ByteGridCoordinate)
	for x := range g.Cols() {
		c := g.ValueAt(x, 0)
		if v, valid := c.Value(); v == 'S' && valid {
			beamXs[x] = true
			break
		}
	}

	for y := 1; y <= g.Cols(); y++ {
		beamClone := slices.Collect(maps.Keys(beamXs))
		for _, x := range beamClone {
			c := g.ValueAt(x, y)
			v, valid := c.Value()
			if !valid {
				continue
			}

			switch v {
			case '^':
				splitters[c.Index()] = c
				delete(beamXs, x)
				cleft := g.ValueAt(c.X()-1, c.Y())
				cright := g.ValueAt(c.X()+1, c.Y())
				if g.SetValueAt(cleft, '|') {
					beamXs[c.X()-1] = true
				}
				if g.SetValueAt(cright, '|') {
					beamXs[c.X()+1] = true
				}
			case '.':
				g.SetValueAt(c, '|')
			}
		}
	}

	fmt.Println(g)

	count := uint64(0)
	for _, s := range splitters {
		val, _ := g.ValueAt(s.X(), s.Y()-1).Value()
		if val == '|' {
			count++
		}
	}

	return count, nil
}

func Day7Part2(in io.Reader) (uint64, error) {
	rows, err := utils.InputAsLines(in, false)
	if err != nil {
		return 0, err
	}

	cols := len(rows[0])
	paths := make([]uint64, cols)

	paths[strings.IndexByte(rows[0], 'S')] = 1
	rows = rows[1:]
	for _, row := range rows {
		for x, r := range row {
			if r == '^' {
				if x > 0 {
					paths[x-1] += paths[x]
				}

				if x < cols-1 {
					paths[x+1] += paths[x]
				}

				paths[x] = 0
			}
		}

		fmt.Println(paths)
	}

	fmt.Println(paths)
	return slices.Reduce(paths, add, 0), nil
}
