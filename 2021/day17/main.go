package main

import (
	"fmt"
	"math"
	"strings"

	"2021/common"
)

type Area struct {
	xRange, yRange common.Coord
}

func NewArea(path string) Area {
	line := common.ReadLines(path)[0]

	split := strings.Split(line[13:], ", ")
	xRangeSplit := strings.Split(split[0][2:], "..")
	yRangeSplit := strings.Split(split[1][2:], "..")

	return Area{
		xRange: [2]int{
			common.StrToInt(xRangeSplit[0]),
			common.StrToInt(xRangeSplit[1]),
		},
		yRange: [2]int{
			common.StrToInt(yRangeSplit[0]),
			common.StrToInt(yRangeSplit[1]),
		},
	}
}

func (area *Area) contains(coord common.Coord) bool {
	return coord.X() <= area.xRange[1] &&
		coord.X() >= area.xRange[0] &&
		coord.Y() >= area.yRange[0] &&
		coord.Y() <= area.yRange[1]
}

func (area *Area) visitAll(visitor func(common.Coord)) {
	minY := int(math.Min(float64(area.yRange[0]), float64(area.yRange[1])))
	maxY := int(math.Max(math.Abs(float64(area.yRange[0])), math.Abs(float64(area.yRange[1]))))
	maxX := int(math.Abs(math.Max(float64(area.xRange[0]), float64(area.xRange[1]))))

	for xv := 1; xv <= maxX; xv++ {
		for yv := minY; yv <= maxY; yv++ {
			visitor([2]int{xv, yv})
		}
	}
}

func interpolate(velocity common.Coord, start common.Coord, steps int) common.Coord {
	s := int(math.Min(float64(velocity.X()), float64(steps)))

	return [2]int{
		start.X() + s*velocity.X() - (s * (s - 1) / 2),
		start.Y() + steps*velocity.Y() - (steps * (steps - 1) / 2),
	}
}

func (area *Area) intersects(velocity common.Coord) (bool, int) {
	maxY := math.MinInt64
	start := [2]int{0, 0}
	step := 1

	for {
		next := interpolate(velocity, start, step)

		if next.Y() > maxY {
			maxY = next.Y()
		}

		if area.contains(next) {
			return true, maxY
		}

		if next.X() > area.xRange[1] || next.Y() < area.yRange[0] {
			break
		}

		step++
	}

	return false, maxY
}

func findMaxY(area Area) int {
	maxY := math.MinInt64

	area.visitAll(func(velocity common.Coord) {
		intr, topY := area.intersects(velocity)

		if intr && topY > maxY {
			maxY = topY
		}
	})

	return maxY
}

func countAll(area Area) int {
	total := 0

	area.visitAll(func(velocity common.Coord) {
		intr, _ := area.intersects(velocity)

		if intr {
			total++
		}
	})

	return total
}

func main() {
	area := NewArea("./.inputs/2021/17/input")

	fmt.Println("Part I", findMaxY(area))
	fmt.Println("Part II", countAll(area))
}
