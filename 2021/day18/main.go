package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"2021/common"
)

type Num struct {
	left, right  int
	lNext, rNext *Num
}

func (num *Num) Str() string {
	var leftStr, rightStr string

	if num.lNext != nil {
		leftStr = num.lNext.Str()
	} else {
		leftStr = fmt.Sprintf("%d", num.left)
	}

	if num.rNext != nil {
		rightStr = num.rNext.Str()
	} else {
		rightStr = fmt.Sprintf("%d", num.right)
	}

	return fmt.Sprintf("[%s,%s]", leftStr, rightStr)
}

func (num *Num) Magnitude() int {
	var lVal, rVal int

	if num.lNext != nil {
		lVal = 3 * num.lNext.Magnitude()
	} else {
		lVal = 3 * num.left
	}

	if num.rNext != nil {
		rVal = 2 * num.rNext.Magnitude()
	} else {
		rVal = 2 * num.right
	}

	return lVal + rVal
}

func splitIx(str string) int {
	depth := 0

	for ix, char := range str {
		switch {
		case char == '[':
			depth++
		case char == ']':
			depth--
		case char == ',' && depth == 1:
			return ix
		}
	}

	panic("splitIx not found")
}

func NewNum(str string) *Num {
	ix := splitIx(str)
	num := Num{}

	leftStr, rightStr := str[1:ix], str[ix+1:len(str)-1]

	value, err := strconv.Atoi(leftStr)
	if err != nil {
		num.lNext = NewNum(leftStr)
	} else {
		num.left = value
	}

	value, err = strconv.Atoi(rightStr)
	if err != nil {
		num.rNext = NewNum(rightStr)
	} else {
		num.right = value
	}

	return &num
}

func add(first *Num, second *Num) *Num {
	return &Num{
		lNext: first,
		rNext: second,
	}
}

func incLeft(num *Num, val int) {
	if num.lNext == nil {
		num.left += val
	} else {
		incLeft(num.lNext, val)
	}
}

func incRight(num *Num, val int) {
	if num.rNext == nil {
		num.right += val
	} else {
		incRight(num.rNext, val)
	}
}

func reduceStep(num *Num, depth int) (bool, bool, int, int) {
	if depth == 4 {
		return true, true, num.left, num.right
	} else {
		if num.lNext != nil {
			terminate, reduced, lVal, rVal := reduceStep(num.lNext, depth+1)

			if reduced {
				num.left = 0
				num.lNext = nil
			}

			if terminate {
				if rVal >= 0 {
					if num.rNext == nil {
						num.right += rVal
					} else {
						incLeft(num.rNext, rVal)
					}

					rVal = -1
				}

				return terminate, false, lVal, rVal
			}
		}

		if num.rNext != nil {
			terminate, reduced, lVal, rVal := reduceStep(num.rNext, depth+1)

			if reduced {
				num.right = 0
				num.rNext = nil
			}

			if terminate {
				if lVal >= 0 {
					if num.lNext == nil {
						num.left += lVal
					} else {
						incRight(num.lNext, lVal)
					}

					lVal = -1
				}

				return terminate, false, lVal, rVal
			}
		}

		return false, false, -1, -1
	}
}

func split(num *Num) bool {
	if num.lNext == nil {
		if num.left >= 10 {
			rem := num.left % 2
			num.lNext = &Num{left: num.left / 2, right: num.left/2 + rem}
			num.left = 0

			return true
		}
	}

	if num.lNext != nil {
		if split(num.lNext) {
			return true
		}
	}

	if num.rNext == nil {
		if num.right >= 10 {
			rem := num.right % 2
			num.rNext = &Num{left: num.right / 2, right: num.right/2 + rem}
			num.right = 0

			return true
		}
	}

	if num.rNext != nil {
		if split(num.rNext) {
			return true
		}
	}

	return false
}

func reduce(num *Num) {
	for {
		reduced, _, _, _ := reduceStep(num, 0)

		if reduced {
			continue
		}

		splitted := split(num)

		if !reduced && !splitted {
			return
		}
	}
}

func main() {
	lines := common.ReadLines(os.Args[1])
	nums := make([]*Num, len(lines))

	for ix, line := range lines {
		nums[ix] = NewNum(line)
	}

	num := nums[0]
	for _, next := range nums[1:] {
		num = add(num, next)
		reduce(num)
	}

	fmt.Println("Part I", num.Magnitude())

	max := math.MinInt64

	for ix, first := range lines {
		for iy, second := range lines {
			if ix == iy {
				continue
			}

			num := add(NewNum(first), NewNum(second))
			reduce(num)
			magnitude := num.Magnitude()
			if magnitude > max {
				max = magnitude
			}
		}
	}

	fmt.Println("Part II", max)
}
