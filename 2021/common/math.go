package common

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
