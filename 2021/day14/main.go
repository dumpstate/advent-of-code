package main

import (
	"fmt"
	"os"
	"strings"

	"2021/common"
)

func parse(path string) (string, map[string]rune) {
	lines := common.ReadLines(path)
	template := lines[0]
	insertionRules := make(map[string]rune)

	for _, line := range lines[2:] {
		split := strings.Split(line, " -> ")
		if len(split) != 2 || len(split[0]) != 2 || len(split[1]) != 1 {
			panic("unexpected input")
		}

		insertionRules[split[0]] = rune(split[1][0])
	}

	return template, insertionRules
}

func asPolymer(str string) *map[string]int {
	polymer := make(map[string]int)

	for i := 0; i < len(str)-1; i++ {
		polymer[str[i:i+2]]++
	}

	return &polymer
}

func grow(polymer *map[string]int, rules *map[string]rune) *map[string]int {
	next := make(map[string]int)

	for pair, count := range *polymer {
		insertion := (*rules)[pair]

		next[string([]rune{rune(pair[0]), insertion})] += count
		next[string([]rune{insertion, rune(pair[1])})] += count
	}

	return &next
}

func score(template string, polymer *map[string]int) int {
	var counter common.Counter = make(map[int]int)
	counter[int(template[len(template)-1])]++

	for key, count := range *polymer {
		counter[int(key[0])] += count
	}

	leastCommon, mostCommon := counter.MinMax()
	return mostCommon - leastCommon
}

func simulate(template string, rules *map[string]rune, steps int) int {
	polymer := asPolymer(template)

	for step := 1; step <= steps; step++ {
		polymer = grow(polymer, rules)
	}

	return score(template, polymer)
}

func main() {
	template, rules := parse(os.Args[1])

	fmt.Println("Part I", simulate(template, &rules, 10))
	fmt.Println("Part II", simulate(template, &rules, 40))
}
