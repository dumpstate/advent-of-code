package main

import (
	"fmt"
	"os"
	"strings"

	"2021/common"
)

type Dice struct {
	size  int
	rolls int
}

func (dice *Dice) roll() int {
	dice.rolls++

	return (dice.rolls-1)%100 + 1
}

func (dice *Dice) play() int {
	return dice.roll() + dice.roll() + dice.roll()
}

func startingPosition(line string) int {
	split := strings.Split(line, ": ")

	return common.StrToInt(split[1])
}

func parse(lines []string) [2]int {
	return [2]int{
		startingPosition(lines[0]),
		startingPosition(lines[1]),
	}
}

func round(dice *Dice, pos *[]int, scores *[]int, target int) {
	for ix := 0; ix < len(*pos); ix++ {
		(*pos)[ix] = ((*pos)[ix]+dice.play()-1)%10 + 1
		(*scores)[ix] += (*pos)[ix]

		if (*scores)[ix] >= target {
			return
		}
	}
}

func playDeterministicDice(pos []int, target int, diceSize int) int {
	scores, board := make([]int, len(pos)), make([]int, len(pos))
	dice := Dice{size: diceSize}
	copy(board, pos)

	for {
		round(&dice, &board, &scores, target)

		if common.Max(scores[:]) >= target {
			break
		}
	}

	return dice.rolls * common.Min(scores[:])
}

type GameState struct {
	pos   [2]int
	score [2]int
}

func (state *GameState) next(player int, throws [3]int) *GameState {
	sum := common.SumAll(throws[:])
	pos := [2]int{state.pos[0], state.pos[1]}
	pos[player] = (pos[player]+sum-1)%10 + 1

	score := [2]int{state.score[0], state.score[1]}
	score[player] += pos[player]

	return &GameState{pos: pos, score: score}
}

func (state *GameState) isDone() bool {
	return common.Max(state.score[:]) >= 21
}

func (state *GameState) winner() int {
	if !state.isDone() {
		return -1
	}

	if state.score[0] > state.score[1] {
		return 0
	}

	return 1
}

type State struct {
	game   GameState
	player int
}

var Cache map[State][2]int = make(map[State][2]int)

func playDiracDice(state State) [2]int {
	wins, exists := Cache[state]
	if exists {
		return wins
	} else {
		wins = [2]int{}
	}

	for _, throws := range common.Cartesian([3]int{1, 2, 3}) {
		next := state.game.next(state.player, throws)
		winner := next.winner()

		if winner == -1 {
			nextState := State{game: *next, player: (state.player + 1) % 2}
			nextWins := playDiracDice(nextState)
			Cache[nextState] = nextWins
			for ix := 0; ix < 2; ix++ {
				wins[ix] += nextWins[ix]
			}
		} else {
			wins[winner]++
		}
	}

	return wins
}

func main() {
	pos := parse(common.ReadLines(os.Args[1]))

	fmt.Println("Part I", playDeterministicDice(pos[:], 1000, 100))

	res := playDiracDice(State{game: GameState{pos: pos}, player: 0})
	fmt.Println("Part II", common.Max(res[:]))
}
