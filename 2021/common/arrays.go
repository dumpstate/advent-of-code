package common

import "log"

func Filter(nums []int, pred func(n int) bool) []int {
	var res []int

	for _, num := range nums {
		if pred(num) {
			res = append(res, num)
		}
	}

	return res
}

func Map(nums []int, fn func(n int) int) []int {
	res := make([]int, len(nums))

	for ix, num := range nums {
		res[ix] = fn(num)
	}

	return res
}

func Slope(nums []int) []int {
	switch {
	case len(nums) == 0:
		log.Fatal("Cannot compute slope of an empty array")
	case len(nums) == 1:
		return make([]int, 0)
	}

	slope := make([]int, len(nums)-1)

	for ix := 1; ix < len(nums); ix++ {
		slope[ix-1] = nums[ix] - nums[ix-1]
	}

	return slope
}

func SubSum(nums []int, start int, length int) int {
	if length == -1 {
		length = len(nums)
	}

	var sum int

	for ix := start; ix < len(nums) && ix < start+length; ix++ {
		sum += nums[ix]
	}

	return sum
}

func SumAll(nums []int) int {
	return SubSum(nums[:], 0, len(nums))
}
