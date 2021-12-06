package main

import (
	"fmt"
	"strings"

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
	lines := common.ReadLines("./.inputs/2021/6/input")
	split := strings.Split(lines[0], ",")
	state := make([]int, len(split))
	for ix, value := range split {
		state[ix] = common.StrToInt(value)
	}

	fmt.Println("Part I", simulate(state, 80))
	fmt.Println("Part II", simulate(state, 256))
}
