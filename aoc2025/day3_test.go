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

const day3ExampleInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestDay3Part1(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day3Part1(strings.NewReader(day3ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(357))
}

func TestDay3Part1Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day3.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	result, err := aoc2025.Day3Part1(bytes.NewReader(data))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}

func TestDay3Part2(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day3Part2(strings.NewReader(day3ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(3121910778619))
}

func TestDay3Part2Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day3.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	result, err := aoc2025.Day3Part2(bytes.NewReader(data))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}
