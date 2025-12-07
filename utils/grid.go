package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Matrix holds a rectangular XxY matrix of values. It can be constructed
// by calling NewMatrix
type Matrix[T any] struct {
	rows     int
	cols     int
	data     []string
	parsed   []*T
	destring func(s string) (T, error)
}

// NewMatrix takes a slice of rows
func NewMatrix[T any](rows []string, destring func(s string) (T, error), rowSplitter func(s string) []string) (*Matrix[T], error) {
	if rowSplitter == nil {
		rowSplitter = strings.Fields
	}

	if destring == nil {
		return nil, errors.New("destring has no default value")
	}

	m := &Matrix[T]{
		rows:     len(rows),
		destring: destring,
	}

	for _, row := range rows {
		cols := rowSplitter(row)
		m.cols = len(cols)
		m.data = append(m.data, cols...)
	}

	m.parsed = make([]*T, len(m.data))
	return m, nil
}

var ErrInvalidPosition = errors.New("Invalid Position (x,y) =  ")

func (m *Matrix[T]) ValueAt(x, y int) (T, error) {
	zeroT := *(new(T))
	idx := y*m.cols + x
	if idx < 0 || idx >= len(m.data) {
		return zeroT, fmt.Errorf("%w = (%d, %d)", ErrInvalidPosition, x, y)
	}

	if m.parsed[idx] != nil {
		return *(m.parsed[idx]), nil
	}

	t, err := m.destring(m.data[idx])
	if err != nil {
		return zeroT, err
	}

	m.parsed[idx] = &t
	return t, nil
}

func (m *Matrix[T]) Rows() int {
	return m.rows
}

func (m *Matrix[T]) Cols() int {
	return m.cols
}
