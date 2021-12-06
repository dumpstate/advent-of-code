package common

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
