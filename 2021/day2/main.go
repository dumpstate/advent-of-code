package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"2021/common"
)

type Command struct {
	direction string
	value     int
}

func navigateSubSimple(cmds []Command) (int, int) {
	var x, y int

	for _, cmd := range cmds {
		switch {
		case cmd.direction == "forward":
			x += cmd.value
		case cmd.direction == "down":
			y += cmd.value
		case cmd.direction == "up":
			y -= cmd.value
		default:
			log.Fatal("Unknown direction", cmd.direction)
		}
	}

	return x, y
}

func navigateSub(cmds []Command) (int, int) {
	var x, y, aim int

	for _, cmd := range cmds {
		switch {
		case cmd.direction == "forward":
			x += cmd.value
			y += aim * cmd.value
		case cmd.direction == "down":
			aim += cmd.value
		case cmd.direction == "up":
			aim -= cmd.value
		default:
			log.Fatal("Unknown direction", cmd.direction)
		}
	}

	return x, y
}

func readCommands(path string) []Command {
	lines := common.ReadLines(path)
	commands := make([]Command, len(lines))

	for ix, line := range lines {
		split := strings.Split(line, " ")
		commands[ix] = Command{split[0], common.StrToInt(split[1])}
	}

	return commands
}

func main() {
	commands := readCommands(os.Args[1])

	x1, y1 := navigateSubSimple(commands[:])
	x2, y2 := navigateSub(commands[:])

	fmt.Println("Part I:", x1*y1)
	fmt.Println("Part II:", x2*y2)
}
