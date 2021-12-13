package main

import (
	"fmt"
	"math"
	"strings"

	"2021/common"
)

type Fold struct {
	Axis  rune
	Index int
}

func NewFold(line string) Fold {
	split := strings.Split(strings.Replace(line, "fold along ", "", 1), "=")
	return Fold{Axis: rune(split[0][0]), Index: common.StrToInt(split[1])}
}

func parseInput(path string) ([]common.Coord, []Fold) {
	lines := common.ReadLines(path)
	dots := make([]common.Coord, 0)
	folds := make([]Fold, 0)

	for _, line := range lines {
		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "fold along"):
			folds = append(folds, NewFold(line))
		default:
			split := strings.Split(line, ",")
			dots = append(dots, [2]int{
				common.StrToInt(split[0]),
				common.StrToInt(split[1]),
			})
		}
	}

	return dots, folds
}

func asMatrix(dots []common.Coord) *common.Matrix {
	maxX, maxY := math.MinInt64, math.MinInt64
	for _, dot := range dots {
		if dot.X() > maxX {
			maxX = dot.X()
		}

		if dot.Y() > maxY {
			maxY = dot.Y()
		}
	}

	matrix := make(common.Matrix, maxY+1)
	for y := 0; y <= maxY; y++ {
		matrix[y] = make([]int, maxX+1)
	}

	for _, dot := range dots {
		matrix[dot.Y()][dot.X()] = 1
	}

	return &matrix
}

func doFold(sheet *common.Matrix, fold Fold) *common.Matrix {
	sizeX, sizeY := sheet.Dimensions()

	switch {
	case fold.Axis == 'y':
		for y := fold.Index + 1; y < sizeY; y++ {
			targetY := fold.Index - (y - fold.Index)
			if targetY < 0 {
				break
			}

			for x := 0; x < sizeX; x++ {
				source := sheet.GetOrPanic([2]int{x, y})
				if source == 1 {
					(*sheet)[targetY][x] = source
				}
			}
		}

		return sheet.TruncatedByY(fold.Index)
	case fold.Axis == 'x':
		for x := fold.Index + 1; x < sizeX; x++ {
			targetX := fold.Index - (x - fold.Index)
			if targetX < 0 {
				break
			}

			for y := 0; y < sizeY; y++ {
				source := sheet.GetOrPanic([2]int{x, y})
				if source == 1 {
					(*sheet)[y][targetX] = source
				}
			}
		}

		return sheet.TruncatedByX(fold.Index)
	default:
		panic("unknown axis")
	}
}

func main() {
	dots, folds := parseInput("./.inputs/2021/13/input")

	sheet := asMatrix(dots[:])
	sheet = doFold(sheet, folds[0])

	fmt.Println("Part I", sheet.Count(func(value int) bool {
		return value == 1
	}))

	for _, fold := range folds[1:] {
		sheet = doFold(sheet, fold)
	}

	fmt.Println("Part II")
	sheet.Print(map[int]string{
		0: " ",
		1: "#",
	})
}
