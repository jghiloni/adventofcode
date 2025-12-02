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
	result, err := aoc2025.Day1Part1(strings.NewReader(day1ExampleInput))
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if result != 3 {
		t.Fatalf("Expected result 3, got %d", result)
	}
}

func TestDay1Part1Real(t *testing.T) {
	_, err := os.Stat("testdata/day1.txt")
	if errors.Is(err, fs.ErrNotExist) {
		return
	}

	b, err := os.ReadFile("testdata/day1.txt")
	if err != nil {
		t.Fatalf("Unexpected err %v", err)
	}

	result, err := aoc2025.Day1Part1(bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("Unexpected err %v", err)
	}

	fmt.Println(result)
}

func TestDay1Part2(t *testing.T) {
	result, err := aoc2025.Day1Part2(strings.NewReader(day1ExampleInput))
	if err != nil {
		t.Fatal("Unexpected err:", err)
	}

	if result != 6 {
		t.Fatalf("Expected result 6, got %d", result)
	}
}

func TestDay1Part2Real(t *testing.T) {
	_, err := os.Stat("testdata/day1.txt")
	if errors.Is(err, fs.ErrNotExist) {
		return
	}

	b, err := os.ReadFile("testdata/day1.txt")
	if err != nil {
		t.Fatalf("Unexpected err %v", err)
	}

	result, err := aoc2025.Day1Part2(bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("Unexpected err %v", err)
	}

	fmt.Println(result)
}
