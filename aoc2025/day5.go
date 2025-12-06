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
	ranges, _, err := collectInput(in)
	if err != nil {
		return 0, err
	}

	normalized := reduceRanges(ranges)
	total := slices.Reduce(normalized, func(subtotal uint64, r utils.NumRange[uint64]) uint64 {
		return subtotal + (r.Max - r.Min + 1)
	}, uint64(0))

	return total, nil
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

func reduceRanges(ranges []utils.NumRange[uint64]) []utils.NumRange[uint64] {
	if len(ranges) == 0 {
		return nil
	}

	reduced := make([]utils.NumRange[uint64], len(ranges))
	copy(reduced, ranges)
	var reduceCount = 0
	for {
		reduced, reduceCount = reduceRangesIteration(reduced)
		if reduceCount == 0 {
			return reduced
		}
	}
}

func reduceRangesIteration(rs []utils.NumRange[uint64]) ([]utils.NumRange[uint64], int) {
	reduced := 0
	slices.SortFunc(rs, func(a, b utils.NumRange[uint64]) int {
		x := int(a.Min) - int(b.Min)
		if x == 0 {
			x = int(a.Max) - int(b.Max)
		}
		return x
	})
	for i := 1; i < len(rs); i++ {
		if rs[i-1].Intersects(rs[i]) {
			rs[i-1] = utils.NumRange[uint64]{Min: min(rs[i-1].Min, rs[i].Min), Max: max(rs[i-1].Max, rs[i].Max)}
			rs = slices.Delete(rs, i, i+1)
			i--
			reduced++
		}
	}

	return rs, reduced
}
