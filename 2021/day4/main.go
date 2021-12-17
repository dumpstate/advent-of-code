package main

import (
	"fmt"
	"os"
	"strings"

	"2021/common"
)

const BOARD_SIZE = 5

type BingoBoard [][]int

func (board *BingoBoard) score(input []int) int {
	lastInput := common.Last(input)
	otherSum := 0

	for x := 0; x < len(*board); x++ {
		for y := 0; y < len((*board)[x]); y++ {
			num := (*board)[x][y]

			if !common.Contains(input, num) {
				otherSum += num
			}
		}
	}

	return otherSum * lastInput
}

func (board *BingoBoard) isSatisfied(input []int) bool {
	for x := 0; x < BOARD_SIZE; x++ {
		var hCount, vCount int

		for y := 0; y < BOARD_SIZE; y++ {
			if common.Contains(input, (*board)[x][y]) {
				hCount++
			}

			if common.Contains(input, (*board)[y][x]) {
				vCount++
			}
		}

		if hCount == BOARD_SIZE || vCount == BOARD_SIZE {
			return true
		}
	}

	return false
}

func bingoBoards(path string) ([]int, []BingoBoard) {
	lines := common.ReadLines(path)

	var input []int
	var bingoBoards []BingoBoard

	for _, numStr := range strings.Split(lines[0], ",") {
		input = append(input, common.StrToInt(numStr))
	}

	var curr []string

	for ix := 2; ix <= len(lines); ix++ {
		if ix == len(lines) || lines[ix] == "" {
			var bingoBoard BingoBoard

			if len(curr) != BOARD_SIZE {
				panic("Unexpected input")
			}

			for _, bingoLine := range curr {
				split := strings.Fields(bingoLine)
				if len(split) != BOARD_SIZE {
					panic("Unexpected input")
				}

				bingoNums := make([]int, len(split))
				for i, num := range split {
					bingoNums[i] = common.StrToInt(num)
				}

				bingoBoard = append(bingoBoard, bingoNums)
			}

			bingoBoards = append(bingoBoards, bingoBoard)
			curr = nil
			continue
		}

		curr = append(curr, lines[ix])
	}

	return input, bingoBoards
}

func execBingo(input []int, boards []BingoBoard) ([]BingoBoard, []int) {
	var winningBoards []BingoBoard
	var winningBoardsIx []int

	for ix, board := range boards {
		if board.isSatisfied(input[:]) {
			winningBoards = append(winningBoards, board)
			winningBoardsIx = append(winningBoardsIx, ix)
		}
	}

	return winningBoards, winningBoardsIx
}

func runBingo(input []int, boards []BingoBoard) int {
	for ix := 1; ix <= len(input); ix++ {
		winningBoards, _ := execBingo(input[0:ix], boards[:])

		if len(winningBoards) == 1 {
			return winningBoards[0].score(input[0:ix])
		}
	}

	panic("Winning board not found")
}

func runReverseBingo(input []int, boards []BingoBoard) int {
	var prevWinningIx []int
	for ix := 1; ix <= len(input); ix++ {
		winningBoards, ixs := execBingo(input[0:ix], boards[:])

		if len(winningBoards) == len(boards) {
			lastWin := common.Head(common.Subtract(ixs[:], prevWinningIx[:]))
			return boards[lastWin].score(input[0:ix])
		}

		prevWinningIx = ixs
	}

	panic("Last winning board not found")
}

func main() {
	input, boards := bingoBoards(os.Args[1])

	fmt.Println("Part I", runBingo(input[:], boards[:]))
	fmt.Println("Part II", runReverseBingo(input[:], boards[:]))
}
