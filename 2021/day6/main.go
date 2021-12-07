package main

import (
	"fmt"

	"2021/common"
)

func simulate(initialState []int, days int) int {
	state := common.NewCounter(initialState)

	for day := 1; day <= days; day++ {
		nextState := common.NewCounter(nil)

		for value, count := range state {
			if value == 0 {
				nextState[6] += count
				nextState[8] += count
			} else {
				nextState[value-1] += count
			}
		}

		state = nextState
	}

	return state.TotalCount()
}

func main() {
	input := common.FirstLineAsInts("./.inputs/2021/6/input")

	fmt.Println("Part I", simulate(input, 80))
	fmt.Println("Part II", simulate(input, 256))
}
