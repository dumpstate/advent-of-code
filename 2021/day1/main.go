package main

import (
	"fmt"

	"2021/common"
)

func CountIncreases(nums []int) int {
	increases := common.Filter(
		common.Map(common.Slope(nums[:]), common.Signum),
		func(num int) bool {
			return num > 0
		})

	return common.SumAll(increases)
}

func part1(nums []int) int {
	return CountIncreases(nums[:])
}

func part2(nums []int) int {
	var tripleSums []int

	for ix := 1; ix < len(nums)-2; ix++ {
		tripleSums = append(tripleSums, common.SubSum(nums[:], ix, 3))
	}

	return CountIncreases(tripleSums[:])
}

func main() {
	nums := common.ReadLinesAsInts("./.inputs/2021/1/input")

	fmt.Println("Part I:", part1(nums[:]))
	fmt.Println("Part II:", part2(nums[:]))
}
