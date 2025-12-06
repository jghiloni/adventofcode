package aoc2025_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	. "github.com/onsi/gomega"
)

const day6ExampleInput = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + 
`

func TestDay6Part1(t *testing.T) {
	RegisterTestingT(t)

	start := time.Now()
	result, err := aoc2025.Day6Part1(strings.NewReader(day6ExampleInput))
	end := time.Now()

	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(4277556))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))

}

func TestDay6Part2(t *testing.T) {
	RegisterTestingT(t)
	start := time.Now()
	result, err := aoc2025.Day6Part2(strings.NewReader(day6ExampleInput))
	end := time.Now()
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(14))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}
