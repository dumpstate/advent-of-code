package main

import (
	"fmt"
	"os"
	"strings"

	"2021/common"
)

type Point struct {
	X, Y int
}

func NewPoint(str string) *Point {
	split := strings.Split(str, ",")
	if len(split) != 2 {
		panic("Unexpected input")
	}

	return &Point{X: common.StrToInt(split[0]), Y: common.StrToInt(split[1])}
}

type Line struct {
	From, To Point
}

func (line *Line) XStep() int {
	return common.Signum(line.To.X - line.From.X)
}

func (line *Line) YStep() int {
	return common.Signum(line.To.Y - line.From.Y)
}

func (line *Line) IsDiagonal() bool {
	return line.From.X != line.To.X && line.From.Y != line.To.Y
}

func (line *Line) Points() []Point {
	var points []Point

	nextPoint := line.From

	for {
		points = append(points, nextPoint)

		if nextPoint == line.To {
			break
		}

		nextPoint = Point{
			X: nextPoint.X + line.XStep(),
			Y: nextPoint.Y + line.YStep(),
		}
	}

	return points
}

func NewLine(str string) *Line {
	split := strings.Split(str, " -> ")
	if len(split) != 2 {
		panic("Unexpected input")
	}

	return &Line{From: *NewPoint(split[0]), To: *NewPoint(split[1])}
}

func countOverlapping(coords []Line, excludeDiagonal bool) int {
	board := make(map[Point]int)

	for _, line := range coords {
		if excludeDiagonal && line.IsDiagonal() {
			continue
		}

		for _, point := range line.Points() {
			board[point]++
		}
	}

	totalOverlapping := 0

	for _, count := range board {
		if count > 1 {
			totalOverlapping++
		}
	}

	return totalOverlapping
}

func main() {
	input := common.ReadLines(os.Args[1])
	lines := make([]Line, len(input))

	for ix, inputStr := range input {
		lines[ix] = *NewLine(inputStr)
	}

	fmt.Println("Part I", countOverlapping(lines[:], true))
	fmt.Println("Part II", countOverlapping(lines[:], false))
}
