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

const day2ExampleInput = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

func TestDay2Part1(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day2Part1(strings.NewReader(day2ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(1227775554))
}

func TestDay2Part1Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day2.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	result, err := aoc2025.Day2Part1(bytes.NewReader(data))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}

func TestDay2Part2(t *testing.T) {
	RegisterTestingT(t)

	result, err := aoc2025.Day2Part2(strings.NewReader(day2ExampleInput))
	Expect(err).NotTo(HaveOccurred())
	Expect(result).To(BeEquivalentTo(4174379265))
}

func TestDay2Part2Real(t *testing.T) {
	RegisterTestingT(t)

	data, err := os.ReadFile("testdata/day2.txt")
	Expect(err == nil || errors.Is(err, fs.ErrNotExist)).To(BeTrue())

	if data == nil {
		return
	}

	result, err := aoc2025.Day2Part2(bytes.NewReader(data))
	Expect(err).NotTo(HaveOccurred())

	fmt.Println(result)
}
