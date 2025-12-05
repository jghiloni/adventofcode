package utils

import (
	"io"
	"strings"

	"github.com/jghiloni/go-commonutils/v3/slices"
)

// ByteGrid is a wrapper around a string that will allow a user
// to get a byte at an (x, y) coordinate, while safely allowing
// addressing invalid coordinates
type ByteGrid struct {
	data string
	rows int
	cols int
}

// ByteGridCoordinate represents a position on a grid
type ByteGridCoordinate struct {
	x, y int
	g    *ByteGrid
}

func (b ByteGridCoordinate) Value() (byte, bool) {
	if b.x < 0 || b.x >= b.g.cols || b.y < 0 || b.y >= b.g.rows {
		return 0, false
	}

	idx := b.y*b.g.cols + b.x
	return b.g.data[idx], true
}

// NewByteGridFromInput works by reading the io.Reader in as a slice of lines,
// calculating the rows and columns, and then joining the lines together
// to a single line without the original newlines
func NewByteGridFromInput(in io.Reader) (*ByteGrid, error) {
	lines, err := InputAsLines(in, false)
	if err != nil {
		return nil, err
	}

	rows := len(lines)
	cols := len(lines[0])

	return &ByteGrid{
		data: strings.Join(lines, ""),
		rows: rows,
		cols: cols,
	}, nil
}

func (g *ByteGrid) Rows() int {
	return g.rows
}

func (g *ByteGrid) Cols() int {
	return g.cols
}

func (g *ByteGrid) ValueAt(x, y int) ByteGridCoordinate {
	return ByteGridCoordinate{
		x: x,
		y: y,
		g: g,
	}
}

func (g *ByteGrid) MatchingNeighbors(home ByteGridCoordinate, matchFunc func(c ByteGridCoordinate) bool) ([]ByteGridCoordinate, int) {
	neighbors := []ByteGridCoordinate{
		g.ValueAt(home.x-1, home.y-1),
		g.ValueAt(home.x, home.y-1),
		g.ValueAt(home.x+1, home.y-1),
		g.ValueAt(home.x-1, home.y),
		g.ValueAt(home.x+1, home.y),
		g.ValueAt(home.x-1, home.y+1),
		g.ValueAt(home.x, home.y+1),
		g.ValueAt(home.x+1, home.y+1),
	}

	matches := slices.Filter(neighbors, matchFunc)
	return matches, len(matches)
}
