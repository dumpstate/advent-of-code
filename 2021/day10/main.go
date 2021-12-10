package main

import (
	"fmt"
	"sort"

	"2021/common"
)

var scoreMap = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScoreMap = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func findIllegalRune(str string) (rune, []rune) {
	var completion []rune

	for _, r := range str {
		switch {
		case r == '(':
			completion = append(completion, ')')
		case r == '[':
			completion = append(completion, ']')
		case r == '{':
			completion = append(completion, '}')
		case r == '<':
			completion = append(completion, '>')
		case len(completion) > 0:
			last := completion[len(completion)-1]

			if last != r {
				return r, completion
			}

			completion = completion[:len(completion)-1]
		case len(completion) == 0:
			return r, completion
		}
	}

	return 0, completion
}

func score(rs []rune) int {
	total := 0

	for _, r := range rs {
		s, exists := completionScoreMap[r]

		if exists {
			total = total*5 + s
		}
	}

	return total
}

func illegalSyntaxScore(lines []string) int {
	total := 0

	for _, line := range lines {
		illegalRune, _ := findIllegalRune(line)
		total += scoreMap[illegalRune]
	}

	return total
}

func totalCompletionScore(lines []string) int {
	var scores []int

	for _, line := range lines {
		illegalRune, completion := findIllegalRune(line)

		if illegalRune != 0 {
			continue
		}

		scores = append(
			scores,
			score(common.ReversedRunes(completion)),
		)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func main() {
	lines := common.ReadLines("./.inputs/2021/10/input")

	fmt.Println("Part I", illegalSyntaxScore(lines[:]))
	fmt.Println("Part II", totalCompletionScore(lines[:]))
}
