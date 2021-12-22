package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"2021/common"
)

type Command struct {
	sign   int
	cuboid Cuboid
}

type Coord struct {
	x, y, z int
}

type Cuboid struct {
	start, end Coord
}

type Reactor struct {
	cuboids map[Cuboid]int
	span    Cuboid
}

func NewCommand(line string) *Command {
	var sign int

	split := strings.Split(line, " ")

	switch {
	case split[0] == "on":
		sign = 1
	case split[0] == "off":
		sign = -1
	default:
		panic("unsupported flag")
	}

	return &Command{
		sign:   sign,
		cuboid: *NewCuboid(split[1]),
	}
}

func NewCuboid(str string) *Cuboid {
	split := strings.Split(str, ",")
	xSplit := strings.Split(split[0][2:], "..")
	ySplit := strings.Split(split[1][2:], "..")
	zSplit := strings.Split(split[2][2:], "..")

	return &Cuboid{
		start: Coord{
			x: common.StrToInt(xSplit[0]),
			y: common.StrToInt(ySplit[0]),
			z: common.StrToInt(zSplit[0]),
		},
		end: Coord{
			x: common.StrToInt(xSplit[1]),
			y: common.StrToInt(ySplit[1]),
			z: common.StrToInt(zSplit[1]),
		},
	}
}

func parse(lines []string) []Command {
	commands := make([]Command, len(lines))

	for ix, line := range lines {
		commands[ix] = *NewCommand(line)
	}

	return commands
}

func (first *Cuboid) intersection(second Cuboid) *Cuboid {
	xMin := common.MaxInt(first.start.x, second.start.x)
	xMax := common.MinInt(first.end.x, second.end.x)
	yMin := common.MaxInt(first.start.y, second.start.y)
	yMax := common.MinInt(first.end.y, second.end.y)
	zMin := common.MaxInt(first.start.z, second.start.z)
	zMax := common.MinInt(first.end.z, second.end.z)

	if xMax >= xMin && yMax >= yMin && zMax >= zMin {
		return &Cuboid{
			start: Coord{x: xMin, y: yMin, z: zMin},
			end:   Coord{x: xMax, y: yMax, z: zMax},
		}
	} else {
		return nil
	}
}

func (cuboid *Cuboid) count() int {
	return (cuboid.end.x - cuboid.start.x + 1) *
		(cuboid.end.y - cuboid.start.y + 1) *
		(cuboid.end.z - cuboid.start.z + 1)
}

func NewReactor(span Cuboid) *Reactor {
	return &Reactor{
		cuboids: make(map[Cuboid]int),
		span:    span,
	}
}

func (reactor *Reactor) count() int {
	total := 0

	for cuboid, count := range reactor.cuboids {
		total += cuboid.count() * count
	}

	return total
}

func (reactor *Reactor) apply(cmd Command) {
	intersections := make(map[Cuboid]int)
	clipped := cmd.cuboid.intersection(reactor.span)

	if clipped == nil {
		return
	}

	for cuboid, currentSign := range reactor.cuboids {
		intersection := cuboid.intersection(*clipped)

		if intersection != nil {
			intersections[*intersection] -= currentSign
		}
	}

	if cmd.sign > 0 {
		intersections[*clipped] += cmd.sign
	}

	for patch, sign := range intersections {
		reactor.cuboids[patch] += sign
	}
}

func (reactor *Reactor) applyAll(commands []Command) {
	for _, command := range commands {
		reactor.apply(command)
	}
}

func main() {
	commands := parse(common.ReadLines(os.Args[1]))
	reactor := NewReactor(
		Cuboid{
			start: Coord{x: -50, y: -50, z: -50},
			end:   Coord{x: 50, y: 50, z: 50},
		},
	)

	reactor.applyAll(commands)

	fmt.Println("Part I", reactor.count())

	unboundedReactor := NewReactor(
		Cuboid{
			start: Coord{x: math.MinInt64, y: math.MinInt64, z: math.MinInt64},
			end:   Coord{x: math.MaxInt64, y: math.MaxInt64, z: math.MaxInt64},
		},
	)

	unboundedReactor.applyAll(commands)

	fmt.Println("Part II", unboundedReactor.count())
}
