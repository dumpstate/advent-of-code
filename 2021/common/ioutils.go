package common

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadLinesAsInts(path string) []int {
	var ints []int
	lines := ReadLines(path)

	for _, line := range lines {
		ints = append(ints, StrToInt(line))
	}

	return ints
}

func FirstLineAsInts(path string) []int {
	lines := ReadLines(path)
	split := strings.Split(lines[0], ",")
	ints := make([]int, len(split))

	for ix, value := range split {
		ints[ix] = StrToInt(value)
	}

	return ints
}
