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
