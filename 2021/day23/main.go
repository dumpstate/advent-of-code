package main

import (
	"fmt"
	"math"
	"os"

	"2021/common"
)

const HallwayLength = 11

var StepCost map[rune]int = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var Amphiods map[int]rune = map[int]rune{
	0: 'A',
	1: 'B',
	2: 'C',
	3: 'D',
}

var Rooms map[rune]int = map[rune]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
}

type Burrow struct {
	hallway [11]rune
	rooms   [4]string
}

func sideRoom(lines []string, index int, length int) string {
	ix := index*2 + 3
	str := ""

	for i := length; i > 0; i-- {
		str += string(lines[i+1][ix])
	}

	return str
}

func NewBurrow(lines []string, roomLength int) *Burrow {
	return &Burrow{
		hallway: [11]rune{
			'.', '.', '.', '.', '.', '.',
			'.', '.', '.', '.', '.',
		},
		rooms: [4]string{
			sideRoom(lines[:], 0, roomLength),
			sideRoom(lines[:], 1, roomLength),
			sideRoom(lines[:], 2, roomLength),
			sideRoom(lines[:], 3, roomLength),
		},
	}
}

func (burrow *Burrow) clone() *Burrow {
	nextBurrow := Burrow{
		hallway: [11]rune{},
		rooms:   [4]string{},
	}

	copy(nextBurrow.hallway[:], burrow.hallway[:])
	copy(nextBurrow.rooms[:], burrow.rooms[:])

	return &nextBurrow
}

func (burrow *Burrow) String() string {
	roomLength := len(burrow.rooms[0])

	str := "\n#############\n"
	str += fmt.Sprintf("#%s#\n", string(burrow.hallway[:]))

	for ix := roomLength - 1; ix >= 0; ix-- {
		str += fmt.Sprintf(
			"###%s#%s#%s#%s###\n",
			string(burrow.rooms[0][ix]),
			string(burrow.rooms[1][ix]),
			string(burrow.rooms[2][ix]),
			string(burrow.rooms[3][ix]),
		)
	}

	str += "  #########"

	return str
}

func (burrow *Burrow) misaligned() [][2]int {
	positions := make([][2]int, 0)
	roomLength := len(burrow.rooms[0])

	for pos := 0; pos < roomLength; pos++ {
		for room := 0; room < 4; room++ {
			if burrow.rooms[room][pos] == '.' {
				continue
			}

			isBlockingMisaligned := false

			for ix := 0; ix < pos; ix++ {
				if rune(burrow.rooms[room][ix]) != Amphiods[room] {
					isBlockingMisaligned = true
					break
				}
			}

			if isBlockingMisaligned ||
				rune(burrow.rooms[room][pos]) != Amphiods[room] {
				positions = append(positions, [2]int{room, pos})
			}
		}
	}

	for pos := 0; pos < HallwayLength; pos++ {
		if burrow.hallway[pos] != '.' {
			positions = append(positions, [2]int{4, pos})
		}
	}

	return positions
}

func (burrow *Burrow) move(source [2]int, target [2]int) *Move {
	nextBurrow, cost := burrow.clone(), 0
	sr, sp, tr, tp := source[0], source[1], target[0], target[1]
	roomLength := len(burrow.rooms[0])

	switch {
	case sr < 4 && sr >= 0:
		amphiod := rune(burrow.rooms[sr][sp])

		switch {
		case tr == sr:
			// nextBurrow.rooms[sr][sp] = '.'
			// nextBurrow.rooms[tr][tp] = amphiod
			nextBurrow.rooms[sr] = common.Update(nextBurrow.rooms[sr], sp, '.')
			nextBurrow.rooms[tr] = common.Update(nextBurrow.rooms[tr], tp, amphiod)
			cost += StepCost[amphiod]
		case tr == 4:
			// nextBurrow.rooms[sr][sp] = '.'
			nextBurrow.rooms[sr] = common.Update(nextBurrow.rooms[sr], sp, '.')
			nextBurrow.hallway[tp] = amphiod

			hallwaySteps := common.Abs(sr*2+2, tp)
			roomSteps := common.Abs(sp, roomLength)
			cost += StepCost[amphiod] * (hallwaySteps + roomSteps)
		default:
			panic("unexpected target room")
		}
	case sr == 4:
		amphiod := burrow.hallway[sp]

		if tr < 0 || tr > 3 {
			panic("unexpected target room")
		}

		nextBurrow.hallway[sp] = '.'
		// nextBurrow.rooms[tr][tp] = amphiod
		nextBurrow.rooms[tr] = common.Update(nextBurrow.rooms[tr], tp, amphiod)

		hallwaySteps := common.Abs(tr*2+2, sp)
		roomSteps := common.Abs(tp, roomLength)
		cost += StepCost[amphiod] * (hallwaySteps + roomSteps)
	default:
		panic("unexpected source room")
	}

	return &Move{burrow: nextBurrow, cost: cost}
}

type Move struct {
	burrow *Burrow
	cost   int
}

func isRoomIndex(ix int) bool {
	return ix == 2 || ix == 4 || ix == 6 || ix == 8
}

func move(burrow *Burrow, source [2]int) []*Move {
	moves := make([]*Move, 0)
	room, pos := source[0], source[1]
	roomLength := len(burrow.rooms[0])

	switch {
	case room < 4 && room >= 0:
		hallwayStart := room*2 + 2
		isBlocked := false

		for ix := pos + 1; ix < roomLength; ix++ {
			if rune(burrow.rooms[room][ix]) != '.' {
				isBlocked = true
				break
			}
		}

		if burrow.hallway[hallwayStart] == '.' && !isBlocked {
			for ix := hallwayStart - 1; ix >= 0; ix-- {
				if isRoomIndex(ix) {
					continue
				}

				if burrow.hallway[ix] != '.' {
					break
				}

				moves = append(moves, burrow.move(source, [2]int{4, ix}))
			}

			for ix := hallwayStart + 1; ix < HallwayLength; ix++ {
				if isRoomIndex(ix) {
					continue
				}

				if burrow.hallway[ix] != '.' {
					break
				}

				moves = append(moves, burrow.move(source, [2]int{4, ix}))
			}
		}
	case room == 4:
		// amphiod is in the hallway
		amphiod := burrow.hallway[pos]
		targetRoom := Rooms[amphiod]
		roomEntry := targetRoom*2 + 2
		onTheWay := false

		if pos < roomEntry {
			for ix := pos + 1; ix <= roomEntry; ix++ {
				if burrow.hallway[ix] != '.' {
					onTheWay = true
				}
			}
		} else {
			for ix := pos - 1; ix >= roomEntry; ix-- {
				if burrow.hallway[ix] != '.' {
					onTheWay = true
				}
			}
		}

		if !onTheWay {
			for ix := 0; ix < roomLength; ix++ {
				r := rune(burrow.rooms[targetRoom][ix])

				if r == '.' {
					moves = append(moves, burrow.move(source, [2]int{targetRoom, ix}))
					break
				}

				if r != '.' && r != amphiod {
					break
				}
			}
		}
	default:
		panic("unexpected room id")
	}

	return moves
}

var Cache map[Burrow]int = make(map[Burrow]int)
var Invalid map[Burrow]bool = make(map[Burrow]bool)

func findMinCost(burrow *Burrow) (int, bool) {
	finalCost, found := Cache[*burrow]
	if found {
		return finalCost, found
	}

	_, foundInvalid := Invalid[*burrow]
	if foundInvalid {
		return 0, false
	}

	positions := burrow.misaligned()

	if len(positions) == 0 {
		return 0, true
	}

	minCost := math.MaxInt64

	for _, pos := range positions {
		nextMoves := move(burrow, pos)

		if len(nextMoves) == 0 {
			continue
		} else {
			for _, nextMove := range nextMoves {
				remCost, foundNextMin := findMinCost(nextMove.burrow)
				if foundNextMin {
					Cache[*nextMove.burrow] = remCost
					moveCost := nextMove.cost + remCost

					if moveCost < minCost {
						minCost = moveCost
						found = true
					}
				} else {
					Invalid[*nextMove.burrow] = true
				}
			}
		}
	}

	return minCost, found
}

func main() {
	lines := common.ReadLines(os.Args[1])
	burrow := NewBurrow(lines[:], 2)
	minCost, _ := findMinCost(burrow)
	fmt.Println("Part I", minCost)

	extraLines := []string{
		"  #D#C#B#A#  ",
		"  #D#B#A#C#  ",
	}
	extLines := make([]string, 3)
	copy(extLines, lines[0:3])
	extLines = append(extLines, extraLines...)
	extLines = append(extLines, lines[3:]...)
	extBurrow := NewBurrow(extLines[:], 4)
	extMinCost, _ := findMinCost(extBurrow)
	fmt.Println("Part II", extMinCost)
}
