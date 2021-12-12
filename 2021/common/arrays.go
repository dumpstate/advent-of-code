package common

import (
	"log"
	"math"
)

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

func Product(nums []int) int {
	product := 1

	for _, value := range nums {
		product *= value
	}

	return product
}

func Subtract(first []int, second []int) []int {
	var diff []int

	for _, x := range first {
		found := false

		for _, y := range second {
			if x == y {
				found = true
				break
			}
		}

		if !found {
			diff = append(diff, x)
		}
	}

	return diff
}

func Contains(nums []int, val int) bool {
	for _, num := range nums {
		if num == val {
			return true
		}
	}

	return false
}

func ContainsStr(arr []string, val string) bool {
	for _, str := range arr {
		if str == val {
			return true
		}
	}

	return false
}

func Head(nums []int) int {
	if len(nums) == 0 {
		log.Fatal("Cannot take head of an empty array")
	}

	return nums[0]
}

func Last(nums []int) int {
	if len(nums) == 0 {
		log.Fatal("Cannot take last element of an empty array")
	}

	return nums[len(nums)-1]
}

func MinMax(nums []int) (int, int) {
	if len(nums) == 0 {
		panic("Cannot compute MinMax of an empty array")
	}

	min, max := math.MaxInt64, math.MinInt64

	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return min, max
}

func ReversedRunes(rs []rune) []rune {
	total := len(rs)
	reversed := make([]rune, total)

	for ix, r := range rs {
		reversed[total-ix-1] = r
	}

	return reversed
}

func FilterStr(arr []string, pred func(string) bool) []string {
	filtered := make([]string, 0)

	for _, str := range arr {
		if pred(str) {
			filtered = append(filtered, str)
		}
	}

	return filtered
}

func MaxCountStr(arr []string) int {
	countMap := make(map[string]int)
	max := math.MinInt64

	for _, str := range arr {
		count := countMap[str] + 1
		if count > max {
			max = count
		}

		countMap[str] = count
	}

	return max
}
