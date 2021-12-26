package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"2021/common"
)

type Block struct {
	xOffset, yOffset, zDiv int
}

func (b *Block) eval(acc int, next int) int {
	var x, z int

	if ((acc % 26) + b.xOffset) == next {
		x = 0
	} else {
		x = 1
	}

	z = acc / b.zDiv
	z = z * (25*x + 1)

	return z + x*(next+b.yOffset)
}

func value(line string) int {
	split := strings.Split(line, " ")

	return common.StrToInt(split[2])
}

func parse(lines []string) []Block {
	blocks := make([]Block, 0)

	for ix, line := range lines {
		if strings.Contains(line, "inp w") {
			blocks = append(
				blocks,
				Block{
					xOffset: value(lines[ix+5]),
					yOffset: value(lines[ix+15]),
					zDiv:    value(lines[ix+4]),
				},
			)
		}
	}

	return blocks
}

var MaxCache map[[2]int]int = make(map[[2]int]int)

func findMax(blocks []Block, ix int, acc int) int {
	res, found := MaxCache[[2]int{ix, acc}]
	if found {
		return res
	}

	for next := 9; next >= 1; next-- {
		nextAcc := blocks[ix].eval(acc, next)

		switch {
		case ix == len(blocks)-1 && nextAcc == 0:
			return next
		case ix < len(blocks)-1:
			res := findMax(blocks[:], ix+1, nextAcc)

			MaxCache[[2]int{ix + 1, nextAcc}] = res

			if res > 0 {
				return next*int(math.Pow(10, float64(13-ix))) + res
			}
		}
	}

	return 0
}

var MinCache map[[2]int]int = make(map[[2]int]int)

func findMin(blocks []Block, ix int, acc int) int {
	res, found := MinCache[[2]int{ix, acc}]
	if found {
		return res
	}

	for next := 1; next <= 9; next++ {
		nextAcc := blocks[ix].eval(acc, next)

		switch {
		case ix == len(blocks)-1 && nextAcc == 0:
			return next
		case ix < len(blocks)-1:
			res := findMin(blocks[:], ix+1, nextAcc)

			MinCache[[2]int{ix + 1, nextAcc}] = res

			if res > 0 {
				return next*int(math.Pow(10, float64(13-ix))) + res
			}
		}
	}

	return 0
}

func main() {
	blocks := parse(common.ReadLines(os.Args[1]))

	fmt.Println("Part I", findMax(blocks[:], 0, 0))
	fmt.Println("Part II", findMin(blocks[:], 0, 0))
}
