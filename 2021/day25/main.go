package main

import (
	"fmt"
	"os"

	"2021/common"
)

type Seafloor [][]rune

func (seafloor *Seafloor) Dimensions() (int, int) {
	return len((*seafloor)[0]), len(*seafloor)
}

func (seafloor *Seafloor) String() string {
	str := "\n"

	for _, row := range *seafloor {
		str += fmt.Sprintf("%s\n", string(row))
	}

	return str
}

func apply(seafloor *Seafloor, moves [][2][2]int) {
	for _, move := range moves {
		(*seafloor)[move[1][1]][move[1][0]] = (*seafloor)[move[0][1]][move[0][0]]
		(*seafloor)[move[0][1]][move[0][0]] = '.'
	}
}

func step(seafloor *Seafloor) int {
	xSize, ySize := seafloor.Dimensions()
	moves := make([][2][2]int, 0)
	totalMoves := 0

	for iy := 0; iy < ySize; iy++ {
		for ix := 0; ix < xSize; ix++ {
			nextIx := (ix + 1) % xSize

			if (*seafloor)[iy][ix] == '>' && (*seafloor)[iy][nextIx] == '.' {
				moves = append(moves, [2][2]int{
					[2]int{ix, iy},
					[2]int{nextIx, iy},
				})
			}
		}
	}

	apply(seafloor, moves[:])
	totalMoves += len(moves)
	moves = make([][2][2]int, 0)

	for ix := 0; ix < xSize; ix++ {
		for iy := 0; iy < ySize; iy++ {
			nextIy := (iy + 1) % ySize

			if (*seafloor)[iy][ix] == 'v' && (*seafloor)[nextIy][ix] == '.' {
				moves = append(moves, [2][2]int{
					[2]int{ix, iy},
					[2]int{ix, nextIy},
				})
			}
		}
	}

	apply(seafloor, moves[:])
	totalMoves += len(moves)

	return totalMoves
}

func main() {
	lines := common.ReadLines(os.Args[1])
	seafloor := make(Seafloor, len(lines))
	for iy, line := range lines {
		seafloor[iy] = make([]rune, len(line))

		for ix, char := range line {
			seafloor[iy][ix] = rune(char)
		}
	}

	count := 1
	for step(&seafloor) > 0 {
		count++
	}

	fmt.Println("Part I", count)
}
