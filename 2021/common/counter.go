package common

import "math"

type Counter map[int]int

func NewCounter(input []int) Counter {
	counter := make(map[int]int)

	for _, value := range input {
		counter[value]++
	}

	return counter
}

func (counter *Counter) TotalCount() int {
	total := 0

	for _, value := range *counter {
		total += value
	}

	return total
}

func (counter *Counter) MinMax() (int, int) {
	min, max := math.MaxInt64, math.MinInt64

	for _, freq := range *counter {
		if freq > max {
			max = freq
		}

		if freq < min {
			min = freq
		}
	}

	return min, max
}
