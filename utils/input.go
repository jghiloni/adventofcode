package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jghiloni/go-commonutils/v3/slices"
)

// InputAsLines reads an io.Reader and breaks it up into a slice of strings
// where each element is a line in the input, with the trailing newline removed.
// If preserveWhitespace is false, strings.TrimSpace is called on the line before
// adding it to the slice.
func InputAsLines(in io.Reader, preserveWhitespace bool) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if !preserveWhitespace {
			line = strings.TrimSpace(line)
		}

		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func InputAsDelimitedScanner(in io.Reader, delimiter string) *bufio.Scanner {
	scanner := bufio.NewScanner(in)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		idx := bytes.Index(data, []byte(delimiter))
		if idx > 0 {
			return idx + 1, data[:idx], nil
		}

		return len(data), data, nil
	})

	return scanner
}

type NumRange[T interface{ NaturalNumber | RealNumber }] struct {
	Min T
	Max T
}

func StringsAsRangeList[T NaturalNumber](list []string) (ranges []NumRange[T], err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			return
		}
	}()

	ranges = slices.Map(list, func(r string) NumRange[T] {
		minmax := strings.Split(r, "-")
		rMin, e := strconv.Atoi(minmax[0])
		if e != nil {
			panic(e)
		}

		rMax, e := strconv.Atoi(minmax[1])
		if e != nil {
			panic(e)
		}

		return NumRange[T]{
			Min: T(rMin),
			Max: T(rMax),
		}
	})

	return
}
