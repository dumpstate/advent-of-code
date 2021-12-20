package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"2021/common"
)

func parse(lines []string) *common.Matrix {
	var mat common.Matrix = make([][]int, len(lines))

	for iy, line := range lines {
		mat[iy] = make([]int, len(line))

		for ix, char := range line {
			if char == '#' {
				mat[iy][ix] = 1
			} else {
				mat[iy][ix] = 0
			}
		}
	}

	return &mat
}

func enhancedIndex(img *common.Matrix, coord common.Coord, off int) int {
	x, y := coord.X(), coord.Y()
	xSize, ySize := img.Dimensions()
	binary := make([]string, 9)

	for iy := y - 1; iy <= y+1; iy++ {
		for ix := x - 1; ix <= x+1; ix++ {
			index := 3*(iy-y+1) + (ix - x + 1)

			if iy < 0 || ix < 0 || iy >= ySize || ix >= xSize {
				binary[index] = fmt.Sprintf("%d", off)
			} else {
				binary[index] = fmt.Sprintf("%d", img.GetOrPanic([2]int{ix, iy}))
			}
		}
	}

	index, err := strconv.ParseInt(strings.Join(binary, ""), 2, 64)
	if err != nil {
		panic("failed to parse binary")
	}

	return int(index)
}

func extend(img *common.Matrix, off int) *common.Matrix {
	xSize, ySize := img.Dimensions()
	var extended common.Matrix = make([][]int, ySize+2)

	for iy := 0; iy < ySize+2; iy++ {
		extended[iy] = make([]int, xSize+2)

		for ix := 0; ix < xSize+2; ix++ {
			if iy == 0 || ix == 0 || iy == ySize+1 || ix == xSize+1 {
				extended[iy][ix] = off
			} else {
				extended[iy][ix] = (*img)[iy-1][ix-1]
			}
		}
	}

	return &extended
}

func enhancedPixel(enhancement string, index int) int {
	enhancedPx := enhancement[index]

	switch {
	case enhancedPx == '#':
		return 1
	case enhancedPx == '.':
		return 0
	default:
		panic("unknown enhanced px")
	}
}

func enhance(enhancement string, img *common.Matrix, off int) (*common.Matrix, int) {
	xSize, ySize := img.Dimensions()
	next := make(map[common.Coord]int)
	enhanced := extend(img, off)

	for iy := 0; iy < ySize+2; iy++ {
		for ix := 0; ix < xSize+2; ix++ {
			coord := [2]int{ix, iy}
			next[coord] = enhancedPixel(enhancement, enhancedIndex(enhanced, coord, off))
		}
	}

	for coord, value := range next {
		(*enhanced)[coord.Y()][coord.X()] = value
	}

	if off == 0 {
		off = enhancedPixel(enhancement, 0)
	} else {
		off = enhancedPixel(enhancement, len(enhancement)-1)
	}

	return enhanced, off
}

func countLitPixels(img *common.Matrix) int {
	total := 0

	for _, row := range *img {
		total += common.SumAll(row)
	}

	return total
}

func main() {
	lines := common.ReadLines(os.Args[1])
	enhancement := lines[0]

	img, off := parse(lines[2:]), 0
	for i := 0; i < 2; i++ {
		img, off = enhance(enhancement, img, off)
	}

	fmt.Println("Part I", countLitPixels(img))

	img, off = parse(lines[2:]), 0
	for i := 0; i < 50; i++ {
		img, off = enhance(enhancement, img, off)
	}

	fmt.Println("Part II", countLitPixels(img))
}
