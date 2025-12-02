package aoc2025

import (
	"fmt"
	"io"
	"maps"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

func Day2Part1(in io.Reader) (uint64, error) {
	scanner := utils.InputAsDelimitedScanner(in, ",")
	var rangeStrings []string

	for scanner.Scan() {
		rangeStrings = append(rangeStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	ranges, err := utils.StringsAsRangeList[uint64](rangeStrings)
	if err != nil {
		return 0, err
	}

	return slices.Reduce(ranges, addSingleRepeatingInvalidIDs, uint64(0)), nil
}

type stringRange struct {
	from       string
	to         string
	invalidIDs []uint64
}

func Day2Part2(in io.Reader) (uint64, error) {
	scanner := utils.InputAsDelimitedScanner(in, ",")
	var rangeStrings []string

	for scanner.Scan() {
		rangeStrings = append(rangeStrings, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	ranges := slices.Map(rangeStrings, func(s string) stringRange {
		p := strings.Split(s, "-")
		r := stringRange{from: p[0], to: p[1]}

		setAllRepeatingInvalidIDs(&r)
		return r
	})

	total := uint64(0)
	for _, r := range ranges {
		for _, i := range r.invalidIDs {
			total += i
		}
	}

	return total, nil
}

func addSingleRepeatingInvalidIDs(total uint64, r utils.NumRange[uint64]) uint64 {
	minDigits := utils.GetNumDigits(r.Min)
	maxDigits := utils.GetNumDigits(r.Max)

	var (
		smallestHalf, smallestRepeater,
		largestHalf, largestRepeater uint64
	)

	switch minDigits % 2 {
	case 0:
		smallestHalf = getSingleRepeater(r.Min)
		smallestRepeater = getSingleInvalidID(smallestHalf)

		if smallestRepeater < r.Min {
			smallestHalf++
			smallestRepeater = getSingleInvalidID(smallestHalf)
		}

	case 1:
		smallestHalf = utils.Base10LeftShift(uint64(1), minDigits/2)
		smallestRepeater = getSingleInvalidID(smallestHalf)
	}

	switch maxDigits % 2 {
	case 0:
		largestHalf = getSingleRepeater(r.Max)
		largestRepeater = getSingleInvalidID(largestHalf)

		if largestRepeater > r.Max {
			largestHalf--
			largestRepeater = getSingleInvalidID(largestRepeater)
		}
	case 1:
		largestRepeater = utils.Base10LeftShift(uint64(1), maxDigits-1) - 1
		largestHalf = getSingleRepeater(largestRepeater)
	}

	subtotal := uint64(0)
	if smallestRepeater > r.Max {
		return total
	}

	for h := smallestHalf; h <= largestHalf; h++ {
		subtotal += getSingleInvalidID(h)
	}

	return total + subtotal
}

func setAllRepeatingInvalidIDs(sr *stringRange) {
	from := sr.from
	minLen := len(sr.from)
	maxLen := len(sr.to)

	if maxLen == 1 {
		return
	}

	if minLen == 1 {
		minLen = 2
		from = "11"
	}

	minv, _ := strconv.ParseUint(from, 10, 0)
	maxv, _ := strconv.ParseUint(sr.to, 10, 0)

	factors := map[int][]int{}
	for i := range maxLen - minLen + 1 {
		x := i + minLen
		factors[x] = getRepeatingLengths(x)
	}

	allRepeatingIDs := map[uint64]bool{}
	for idLen, lenFactors := range factors {
		for _, factor := range lenFactors {
			segmentMin := utils.Base10LeftShift(uint64(1), factor-1)
			segmentMaxExcl := utils.Base10LeftShift(uint64(1), factor)

			for r := segmentMin; r < segmentMaxExcl; r++ {
				rid := getRepeatingID(fmt.Sprintf("%d", r), idLen)
				if rid < minv {
					continue
				}

				if rid > maxv {
					break
				}
				allRepeatingIDs[rid] = true
			}
		}
	}

	sr.invalidIDs = slices.Collect(maps.Keys(allRepeatingIDs))
}

func getSingleRepeater(full uint64) uint64 {
	digits := utils.GetNumDigits(full)
	halfDigits := digits / 2
	if digits%2 == 1 {
		halfDigits++
	}
	return full / utils.Base10LeftShift(uint64(1), halfDigits)
}

func getSingleInvalidID(half uint64) uint64 {
	digits := utils.GetNumDigits(half)
	return utils.Base10LeftShift(half, digits) + half
}

func getRepeatingLengths(len int) []int {
	factors := []int{1}

	for i := 2; i <= len/2; i++ {
		if len%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

func getRepeatingID(repeat string, l int) uint64 {
	idStr := strings.Repeat(repeat, l/len(repeat))
	id, _ := strconv.ParseUint(idStr, 10, 0)
	return id
}
