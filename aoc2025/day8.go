package aoc2025

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"reflect"
	"strconv"
	"strings"

	"github.com/jghiloni/adventofcode/utils"
	"github.com/jghiloni/go-commonutils/v3/slices"
)

type junctionBox struct {
	x   int
	y   int
	z   int
	od2 int
	xyz string
}

type circuit struct {
	boxes map[string]bool
}

type connection struct {
	j1  *junctionBox
	j2  *junctionBox
	dsq int
}

func Day8Part1(in io.Reader) (uint64, error) {
	lines, err := utils.InputAsLines(in, false)
	if err != nil {
		return 0, err
	}

	count, err := stringToUint(lines[0])
	if err != nil {
		return 0, err
	}

	topN, err := stringToUint(lines[1])
	if err != nil {
		return 0, err
	}

	lines = lines[2:]
	allBoxes := make([]*junctionBox, len(lines))
	for i := range lines {
		allBoxes[i], err = newJunctionBox(lines[i])
		if err != nil {
			return 0, err
		}
	}

	slices.SortFunc(allBoxes, func(a, b *junctionBox) int {
		diff := a.od2 - b.od2
		if diff == 0 {
			diff = a.x - b.x
		}

		if diff == 0 {
			diff = a.y - b.y
		}

		if diff == 0 {
			diff = a.z - b.z
		}

		return diff
	})

	var connections []*connection
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			connections = append(connections, &connection{
				j1:  allBoxes[i],
				j2:  allBoxes[j],
				dsq: dsquared(allBoxes[i], allBoxes[j]),
			})
		}
	}

	slices.SortFunc(connections, func(a, b *connection) int {
		return a.dsq - b.dsq
	})

	connections = connections[:count]
	var circuits []*circuit

	for _, p := range connections {
		matchingCircuits := slices.Filter(circuits, func(c *circuit) bool {
			return c.contains(p.j1, p.j2)
		})

		switch len(matchingCircuits) {
		case 0:
			circuits = append(circuits, newCircuit(p.j1, p.j2))
		case 1:
			idx := slices.IndexFunc(circuits, func(c *circuit) bool {
				return c.contains(p.j1, p.j2)
			})
			circuits[idx].boxes[p.j1.xyz] = true
			circuits[idx].boxes[p.j2.xyz] = true
		default:
			newc := newCircuit()
			for _, m := range matchingCircuits {
				for k := range maps.Keys(m.boxes) {
					newc.boxes[k] = true
				}

				idx := slices.IndexFunc(circuits, func(c *circuit) bool {
					return c.equals(m)
				})

				circuits = append(circuits[:idx], circuits[idx+1:]...)
			}
			circuits = append(circuits, newc)
		}
	}

	slices.SortFunc(circuits, func(a, b *circuit) int {
		return len(b.boxes) - len(a.boxes)
	})

	return slices.Reduce(slices.Map(circuits[:topN], func(c *circuit) uint64 {
		return uint64(len(c.boxes))
	}), mul, 1), nil
}

func Day8Part2(in io.Reader) (uint64, error) {
	lines, err := utils.InputAsLines(in, false)
	if err != nil {
		return 0, err
	}

	lines = lines[2:]
	allBoxes := make([]*junctionBox, len(lines))
	circuits := make([]*circuit, len(lines))
	for i := range lines {
		allBoxes[i], err = newJunctionBox(lines[i])
		if err != nil {
			return 0, err
		}
		circuits[i] = newCircuit(allBoxes[i])
	}

	slices.SortFunc(allBoxes, func(a, b *junctionBox) int {
		diff := a.od2 - b.od2
		if diff == 0 {
			diff = a.x - b.x
		}

		if diff == 0 {
			diff = a.y - b.y
		}

		if diff == 0 {
			diff = a.z - b.z
		}

		return diff
	})

	var connections []*connection
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			connections = append(connections, &connection{
				j1:  allBoxes[i],
				j2:  allBoxes[j],
				dsq: dsquared(allBoxes[i], allBoxes[j]),
			})
		}
	}

	slices.SortFunc(connections, func(a, b *connection) int {
		return a.dsq - b.dsq
	})

	var lastConnection *connection
	for i, p := range connections {
		matchingCircuits := slices.Filter(circuits, func(c *circuit) bool {
			return c.contains(p.j1, p.j2)
		})

		newc := newCircuit()
		for _, m := range matchingCircuits {
			for k := range maps.Keys(m.boxes) {
				newc.boxes[k] = true
			}

			idx := slices.IndexFunc(circuits, func(c *circuit) bool {
				return c.equals(m)
			})

			circuits = append(circuits[:idx], circuits[idx+1:]...)
		}
		circuits = append(circuits, newc)

		if len(circuits) == 1 {
			fmt.Printf("converged after %d of %d connections\n", i+1, len(connections))
			lastConnection = p
			break
		}
	}

	if lastConnection == nil {
		return 0, errors.New("never converged")
	}

	return uint64(lastConnection.j1.x * lastConnection.j2.x), nil
}

func newJunctionBox(xyz string) (*junctionBox, error) {
	coords := strings.Split(xyz, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(coords[1])
	if err != nil {
		return nil, err
	}

	z, err := strconv.Atoi(coords[2])
	if err != nil {
		return nil, err
	}

	od2 := x*x + y*y + z*z

	return &junctionBox{x, y, z, od2, xyz}, nil
}

func newCircuit(jbs ...*junctionBox) *circuit {
	c := &circuit{boxes: make(map[string]bool)}
	for _, jb := range jbs {
		c.boxes[jb.xyz] = true
	}

	return c
}

func (c *circuit) contains(js ...*junctionBox) bool {
	for _, j := range js {
		if c.boxes[j.xyz] {
			return true
		}
	}

	return false
}

func (c *circuit) equals(c2 *circuit) bool {
	return reflect.DeepEqual(c.boxes, c2.boxes)
}

func dsquared(j1, j2 *junctionBox) int {
	dsq := (j1.x-j2.x)*(j1.x-j2.x) +
		(j1.y-j2.y)*(j1.y-j2.y) +
		(j1.z-j2.z)*(j1.z-j2.z)

	return dsq
}
