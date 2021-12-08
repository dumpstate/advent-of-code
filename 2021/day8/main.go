package main

import (
	"fmt"
	"math"
	"strings"

	"2021/common"
)

type Entry struct {
	Patterns []string
	Digits   []string
}

func NewEntry(input string) Entry {
	split := strings.Split(input, " | ")

	return Entry{
		Patterns: strings.Fields(split[0]),
		Digits:   strings.Fields(split[1]),
	}
}

func countEasyDigits(entries []Entry) int {
	count := 0

	for _, entry := range entries {
		for _, digit := range entry.Digits {
			l := len(digit)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				count++
			}
		}
	}

	return count
}

func ofLen(pattern []string, length int) []string {
	var res []string

	for _, digit := range pattern {
		if len(digit) == length {
			res = append(res, digit)
		}
	}

	return res
}

func match(pattern []string, length int, contains []string, doesNotContain []string) string {
	for _, lenMatch := range ofLen(pattern[:], length) {
		totalContains, totalNotContains := 0, 0

		for _, str := range contains {
			if common.ContainsInAnyOrder(lenMatch, str) {
				totalContains++
			}
		}

		for _, str := range doesNotContain {
			if !common.ContainsInAnyOrder(lenMatch, str) {
				totalNotContains++
			}
		}

		if len(contains) == totalContains && len(doesNotContain) == totalNotContains {
			return lenMatch
		}
	}

	panic("Not found")
}

func numMapping(pattern []string) map[int]string {
	mapping := make(map[int]string)

	mapping[1] = match(pattern[:], 2, nil, nil)
	mapping[4] = match(pattern[:], 4, nil, nil)
	mapping[7] = match(pattern[:], 3, nil, nil)
	mapping[8] = match(pattern[:], 7, nil, nil)
	mapping[3] = match(pattern[:], 5, []string{mapping[1]}, nil)
	mapping[9] = match(pattern[:], 6, []string{mapping[4]}, nil)

	btmLeft := common.SubtractStr(mapping[8], mapping[9])
	topLeft := common.SubtractStr(mapping[4], mapping[3])
	middle := common.SubtractStr(common.IntersectStr(mapping[3], mapping[4]), mapping[1])

	mapping[6] = match(pattern[:], 6, []string{btmLeft, middle}, nil)
	mapping[2] = match(pattern[:], 5, []string{btmLeft}, nil)
	mapping[5] = match(pattern[:], 5, []string{topLeft}, nil)
	mapping[0] = match(pattern[:], 6, []string{topLeft}, []string{middle})

	return mapping
}

func find(mapping map[int]string, digit string) int {
	for num, pattern := range mapping {
		if len(pattern) == len(digit) && common.ContainsInAnyOrder(pattern, digit) {
			return num
		}
	}

	panic("Digit does not have a match")
}

func translate(mapping map[int]string, digits []string) int {
	total := 0.0

	for ix, digit := range digits {
		total += math.Pow(10.0, float64(len(digits)-ix-1)) * float64(find(mapping, digit))
	}

	return int(total)
}

func deduceAll(entries []Entry) int {
	total := 0

	for _, entry := range entries {
		total += translate(numMapping(entry.Patterns[:]), entry.Digits)
	}

	return total
}

func main() {
	lines := common.ReadLines("./.inputs/2021/8/input")
	entries := make([]Entry, len(lines))
	for ix, line := range lines {
		entries[ix] = NewEntry(line)
	}

	fmt.Println("Part I", countEasyDigits(entries[:]))
	fmt.Println("Part II", deduceAll(entries[:]))
}
