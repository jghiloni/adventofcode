package aoc2025_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	. "github.com/onsi/gomega"
)

const day5ExampleInput = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

func TestDay5Part1(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day5Part1(strings.NewReader(day5ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(3))
}

func TestDay5Part2(t *testing.T) {
	RegisterTestingT(t)
	start := time.Now()
	result, err := aoc2025.Day5Part2(strings.NewReader(day5ExampleInput))
	end := time.Now()
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(14))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}
