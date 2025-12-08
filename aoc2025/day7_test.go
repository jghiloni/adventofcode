package aoc2025_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	. "github.com/onsi/gomega"
)

const day7ExampleInput = `
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`

func TestDay7Part1(t *testing.T) {
	RegisterTestingT(t)

	start := time.Now()
	result, err := aoc2025.Day7Part1(strings.NewReader(day7ExampleInput))
	end := time.Now()

	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(21))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}

func TestDay7Part2(t *testing.T) {
	RegisterTestingT(t)

	start := time.Now()
	result, err := aoc2025.Day7Part2(strings.NewReader(day7ExampleInput))
	end := time.Now()

	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(40))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}
