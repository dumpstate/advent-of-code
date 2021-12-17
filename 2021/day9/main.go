package main

import (
	"fmt"
	"os"
	"sort"

	"2021/common"
)

func isLowPoint(mat common.Matrix, coord common.Coord) bool {
	point, _ := mat.Get(coord)

	for _, neighbor := range mat.Neighbors(coord) {
		if point >= mat.GetOrPanic(neighbor) {
			return false
		}
	}

	return true
}

func basinSize(points common.Matrix, coord common.Coord) int {
	basin := map[common.Coord]bool{coord: true}
	toVisit := basin

	for {
		if len(toVisit) == 0 {
			break
		}

		nextToVisit := make(map[common.Coord]bool)

		for point := range toVisit {
			for _, neighbor := range points.Neighbors(point) {
				_, contains := basin[neighbor]
				if points.GetOrPanic(neighbor) < 9 && !contains {
					nextToVisit[neighbor] = true
					basin[neighbor] = true
				}
			}
		}

		toVisit = nextToVisit
	}

	return len(basin)
}

func riskLevel(points common.Matrix) int {
	var lowPoints []int

	for y, vector := range points {
		for x, point := range vector {
			if isLowPoint(points[:], [2]int{x, y}) {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return common.SumAll(lowPoints) + len(lowPoints)
}

func targetBasinsProduct(points common.Matrix) int {
	var basinSizes []int

	for y, vector := range points {
		for x := range vector {
			if isLowPoint(points[:], [2]int{x, y}) {
				basinSizes = append(basinSizes, basinSize(points[:], [2]int{x, y}))
			}
		}
	}

	sort.Ints(basinSizes[:])

	return common.Product(basinSizes[len(basinSizes)-3:])
}

func main() {
	points := common.NewMatrix(os.Args[1], "")

	fmt.Println("Part I", riskLevel(points[:]))
	fmt.Println("Part II", targetBasinsProduct(points[:]))
}
