package aoc2025

import (
	"io"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

func Day5Part1(in io.Reader) (uint64, error) {
	ranges, ids, err := collectInput(in)
	if err != nil {
		return 0, err
	}

	slices.SortFunc(ranges, func(a, b utils.NumRange[uint64]) int {
		diff := int(a.Min) - int(b.Min)
		if diff == 0 {
			// if the min is the same, then sort in reverse order by max (implying that larger ranges should be searched first)
			diff = int(b.Max) - int(a.Max)
		}
		return diff
	})

	total := uint64(0)
	for _, id := range ids {
		for _, r := range ranges {
			if r.Min <= id && id <= r.Max {
				total++
				break
			}
		}
	}

	return total, nil
}

func Day5Part2(in io.Reader) (uint64, error) {
	panic("unimplemented")
}

func collectInput(in io.Reader) ([]utils.NumRange[uint64], []uint64, error) {
	lines, err := utils.InputAsLines(in, false)
	if err != nil {
		return nil, nil, err
	}

	rangeStrings := slices.SubsliceUntil(lines, func(item string) bool {
		return strings.IndexByte(item, '-') < 0
	})

	slices.Reverse(lines)
	idStrings := slices.SubsliceUntil(lines, func(item string) bool {
		return strings.IndexByte(item, '-') > 0
	})

	ranges, err := utils.StringsAsRangeList[uint64](rangeStrings)
	if err != nil {
		return nil, nil, err
	}

	ids := slices.Map(idStrings, func(idStr string) uint64 {
		v, _ := strconv.ParseUint(idStr, 10, 0)
		return v
	})

	return ranges, ids, nil
}
