package aoc2025_test

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/jghiloni/adventofcode/aoc2025"

	. "github.com/onsi/gomega"
)

const day1ExampleInput = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func TestDay1Part1(t *testing.T) {
	RegisterTestingT(t)
	result, err := aoc2025.Day1Part1(strings.NewReader(day1ExampleInput))
	Expect(err).ShouldNot(HaveOccurred())
	Expect(result).To(Equal(3))
}

func TestDay1Part1Real(t *testing.T) {
	RegisterTestingT(t)

	_, err := os.Stat("testdata/day1.txt")
	if errors.Is(err, fs.ErrNotExist) {
		return
	}
	Expect(err).NotTo(HaveOccurred())

	b, err := os.ReadFile("testdata/day1.txt")
	Expect(err).NotTo(HaveOccurred())

	result, err := aoc2025.Day1Part1(bytes.NewBuffer(b))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}

func TestDay1Part2(t *testing.T) {
	RegisterTestingT(t)
	result, err := aoc2025.Day1Part2(strings.NewReader(day1ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(Equal(6))
}

func TestDay1Part2Real(t *testing.T) {
	RegisterTestingT(t)
	_, err := os.Stat("testdata/day1.txt")
	if errors.Is(err, fs.ErrNotExist) {
		return
	}
	Expect(err).NotTo(HaveOccurred())

	b, err := os.ReadFile("testdata/day1.txt")
	Expect(err).NotTo(HaveOccurred())

	result, err := aoc2025.Day1Part2(bytes.NewBuffer(b))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}
