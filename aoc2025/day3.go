package aoc2025

import (
	"cmp"
	"io"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

func Day3Part1(in io.Reader) (uint64, error) {
	banks, err := utils.InputAsLines(in, false)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	for _, bank := range banks {
		chargiest, idx := getJoltiestBattery(bank[0 : len(bank)-1])
		nextChargiest, _ := getJoltiestBattery(bank[idx+1:])

		total += uint64(chargiest*10 + nextChargiest)
	}

	return total, nil
}

func Day3Part2(in io.Reader) (uint64, error) {
	banks, err := utils.InputAsLines(in, false)
	if err != nil {
		return 0, err
	}

	total := uint64(0)
	for _, bank := range banks {
		subtotal, err := getMaxCharge(bank, 12)
		if err != nil {
			return 0, err
		}
		total += subtotal
	}

	return total, nil
}

type indexedByte struct {
	index int
	value byte
}

func getMaxCharge(bank string, activateCount int) (uint64, error) {
	chargeBytes := make([]byte, 0, activateCount)
	start := 0
	for endOffset := activateCount - 1; endOffset >= 0; endOffset-- {
		subBank := bank[start : len(bank)-endOffset]
		maxVal := slices.Max([]byte(subBank))
		idx := strings.IndexByte(subBank, maxVal)
		chargeBytes = append(chargeBytes, maxVal)
		start += idx + 1
	}

	charge := string(chargeBytes)
	return strconv.ParseUint(charge, 10, 0)
}

func getJoltiestBattery(bank string) (int, int) {
	batteries := []byte(bank)
	chargiest := variadicMax(batteries...)
	idx := len(bank) - 1
	for ; idx >= 0; idx-- {
		if batteries[idx] == chargiest {
			break
		}
	}

	return int(chargiest - '0'), idx
}

func variadicMax[T cmp.Ordered](args ...T) T {
	if len(args) == 0 {
		panic("no arguments")
	}

	m := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > m {
			m = args[i]
		}
	}

	return m
}

func onlyStragglersRemain(s string) bool {
	s = strings.ReplaceAll(s, "0", "")
	if len(s) == 0 {
		return true
	}

	r := strings.Repeat(s[0:1], len(s))
	return s == r
}
