package aoc2025_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	. "github.com/onsi/gomega"
)

const day4ExampleInput = `
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`

func TestDay4Part1(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day4Part1(strings.NewReader(day4ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(13))
}

func TestDay4Part1Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day4.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	result, err := aoc2025.Day4Part1(bytes.NewReader(data))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}

func TestDay4Part2(t *testing.T) {
	RegisterTestingT(t)
	start := time.Now()
	result, err := aoc2025.Day4Part2(strings.NewReader(day4ExampleInput))
	end := time.Now()
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(43))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}

func TestDay4Part2Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day4.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	start := time.Now()
	result, err := aoc2025.Day4Part2(bytes.NewReader(data))
	end := time.Now()
	Expect(err).NotTo(HaveOccurred())

	fmt.Printf("Elapsed time: %s\n", end.Sub(start))
	fmt.Println(result)
}
