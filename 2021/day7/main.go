package main

import (
	"fmt"
	"math"
	"os"

	"2021/common"
)

func fuelCost(num int, pos int) int {
	cost := 0
	distance := common.Abs(num, pos)

	for ix := 1; ix <= distance; ix++ {
		cost += ix
	}

	return cost
}

func minFuelCost(positions []int, costFn func(n int, p int) int) int {
	min, max := common.MinMax(positions[:])
	minFuelCost := math.MaxInt64

	for targetPos := min; targetPos <= max; targetPos++ {
		cost := 0

		for _, pos := range positions {
			cost += costFn(pos, targetPos)
		}

		if cost < minFuelCost {
			minFuelCost = cost
		}
	}

	return minFuelCost
}

func main() {
	input := common.FirstLineAsInts(os.Args[1])

	fmt.Println("Part I", minFuelCost(input[:], common.Abs))
	fmt.Println("Part II", minFuelCost(input[:], fuelCost))
}
