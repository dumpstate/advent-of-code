package common

import (
	"errors"
	"strings"
)

type Coord [2]int

func (coord *Coord) X() int {
	return (*coord)[0]
}

func (coord *Coord) Y() int {
	return (*coord)[1]
}

type Matrix [][]int

func NewMatrix(path string, rowSep string) Matrix {
	lines := ReadLines(path)
	matrix := make([][]int, len(lines))

	for y, line := range lines {
		matrix[y] = make([]int, len(line))

		for x, r := range strings.Split(line, rowSep) {
			matrix[y][x] = StrToInt(r)
		}
	}

	return matrix
}

func (mat *Matrix) Get(coord Coord) (int, error) {
	if coord.X() < 0 ||
		coord.Y() < 0 ||
		coord.Y() >= len(*mat) ||
		coord.X() >= len((*mat)[coord.Y()]) {
		return 0, errors.New("out of bounds")
	}

	return (*mat)[coord.Y()][coord.X()], nil
}

func (mat *Matrix) Find(pred func(coord Coord) bool) (Coord, bool) {
	xDim, yDim := mat.Dimensions()

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			coord := [2]int{x, y}
			if pred(coord) {
				return coord, true
			}
		}
	}

	return [2]int{}, false
}

func (mat *Matrix) VisitAll(fn func(value *int)) {
	xDim, yDim := mat.Dimensions()

	for x := 0; x < xDim; x++ {
		for y := 0; y < yDim; y++ {
			fn(&(*mat)[y][x])
		}
	}
}

func (mat *Matrix) GetOrPanic(coord Coord) int {
	value, err := (*mat).Get(coord)

	if err != nil {
		panic(err)
	}

	return value
}

func (mat *Matrix) Neighbors(coord Coord) []Coord {
	var neighbors []Coord
	x, y := coord.X(), coord.Y()

	for _, point := range []Coord{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	} {
		_, err := (*mat).Get(point)
		if err == nil {
			neighbors = append(neighbors, point)
		}
	}

	return neighbors
}

func (mat *Matrix) Dimensions() (int, int) {
	return len((*mat)[0]), len(*mat)
}

func (mat *Matrix) TotalEntries() int {
	xDim, yDim := mat.Dimensions()

	return xDim * yDim
}