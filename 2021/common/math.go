package common

import "math"

func Signum(num int) int {
	switch {
	case num < 0:
		return -1
	case num > 0:
		return 1
	}

	return 0
}

func Abs(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}

	return num2 - num1
}

func MinInt(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func MaxInt(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
