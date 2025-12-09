package aoc2025_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/jghiloni/adventofcode/aoc2025"
	. "github.com/onsi/gomega"
)

const day8ExampleInput = `10
3
162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

func TestDay8Part1(t *testing.T) {
	RegisterTestingT(t)

	start := time.Now()
	result, err := aoc2025.Day8Part1(strings.NewReader(day8ExampleInput))
	end := time.Now()

	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(40))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}

func TestDay8Part2(t *testing.T) {
	RegisterTestingT(t)

	start := time.Now()
	result, err := aoc2025.Day8Part2(strings.NewReader(day8ExampleInput))
	end := time.Now()

	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(25272))
	fmt.Printf("Elapsed time on example input: %s\n", end.Sub(start))
}
