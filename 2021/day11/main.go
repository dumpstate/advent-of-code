package main

import (
	"fmt"

	"2021/common"
)

func bumpNeighbors(octopuses common.Matrix, coord common.Coord) {
	xDim, yDim := octopuses.Dimensions()

	for x := coord.X() - 1; x <= coord.X()+1; x++ {
		for y := coord.Y() - 1; y <= coord.Y()+1; y++ {
			if (x == coord.X() && y == coord.Y()) ||
				x < 0 || x >= xDim ||
				y < 0 || y >= yDim {
				continue
			}

			octopuses[y][x]++
		}
	}
}

func flash(octopuses common.Matrix) *map[common.Coord]bool {
	flashed := make(map[common.Coord]bool)

	octopuses.VisitAll(func(octopus *int) {
		(*octopus)++
	})

	for {
		next, found := octopuses.Find(func(coord common.Coord) bool {
			_, contains := flashed[coord]
			return octopuses.GetOrPanic(coord) > 9 && !contains
		})

		if !found {
			break
		}

		bumpNeighbors(octopuses[:], next)
		flashed[next] = true
	}

	for coord := range flashed {
		octopuses[coord.Y()][coord.X()] = 0
	}

	return &flashed
}

func countFlashes(octopuses common.Matrix, steps int) int {
	totalFlashes := 0

	for step := 1; step <= steps; step++ {
		flashed := flash(octopuses[:])

		totalFlashes += len(*flashed)
	}

	return totalFlashes
}

func simultaneousFlash(octopuses common.Matrix) int {
	totalOctopuses := octopuses.TotalEntries()

	for step := 1; true; step++ {
		flashed := flash(octopuses[:])

		if len(*flashed) == totalOctopuses {
			return step
		}
	}

	panic("they never synchronize")
}

func main() {
	path := "./.inputs/2021/11/input"
	octopuses1 := common.NewMatrix(path, "")
	octopuses2 := common.NewMatrix(path, "")

	fmt.Println("Part I", countFlashes(octopuses1[:], 100))
	fmt.Println("Part II", simultaneousFlash(octopuses2[:]))
}
