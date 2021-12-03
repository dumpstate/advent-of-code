package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"2021/common"
)

func binToInt(binaryNum string) int64 {
	value, err := strconv.ParseInt(binaryNum, 2, 64)
	if err != nil {
		log.Fatal("Failed to convert to decimal", binaryNum)
	}

	return value
}

func bitArrToInt(binaryNum []string) int64 {
	return binToInt(strings.Join(binaryNum, ""))
}

func countBits(binaryNums []string, pos int) (int, int) {
	var ones, zeroes int

	for _, num := range binaryNums {
		if string(num[pos]) == "1" {
			ones++
		} else {
			zeroes++
		}
	}

	return ones, zeroes
}

func mostCommonBit(lines []string, pos int) string {
	ones, zeroes := countBits(lines[:], pos)

	if ones >= zeroes {
		return "1"
	}

	return "0"
}

func leastCommonBit(lines []string, pos int) string {
	ones, zeroes := countBits(lines[:], pos)

	if ones >= zeroes {
		return "0"
	}

	return "1"
}

func powerConsumption(lines []string) int64 {
	length := len(lines[0])
	counts := make([]int, length)
	gammaBits := make([]string, length)
	epsilonBits := make([]string, length)

	for _, line := range lines {
		for ix, r := range line {
			if r == '1' {
				counts[ix]++
			}
		}
	}

	for ix := 0; ix < length; ix++ {
		if counts[ix] > len(lines)/2 {
			gammaBits[ix] = "1"
			epsilonBits[ix] = "0"
		} else {
			gammaBits[ix] = "0"
			epsilonBits[ix] = "1"
		}
	}

	return bitArrToInt(gammaBits) * bitArrToInt(epsilonBits)
}

func filteredRating(binaryNums []string, criteria func(ns []string, ix int) string) int64 {
	current := binaryNums

	for ix := 0; ix < len(binaryNums[0]); ix++ {
		bit := criteria(current[:], ix)

		var filtered []string
		for _, num := range current {
			if string(num[ix]) == bit {
				filtered = append(filtered, num)
			}
		}

		switch {
		case len(filtered) == 1:
			return binToInt(filtered[0])
		case len(filtered) == 0:
			log.Fatal("Empty filtered array")
		default:
			current = filtered
			continue
		}
	}

	return 0
}

func oxygenGeneratorRating(binaryNums []string) int64 {
	return filteredRating(binaryNums[:], mostCommonBit)
}

func co2ScrubberRating(binaryNums []string) int64 {
	return filteredRating(binaryNums[:], leastCommonBit)
}

func lifeSupportRating(lines []string) int64 {
	return oxygenGeneratorRating(lines[:]) * co2ScrubberRating(lines[:])
}

func main() {
	binaryNums := common.ReadLines("./.inputs/2021/3/input")

	fmt.Println("Part I:", powerConsumption(binaryNums[:]))
	fmt.Println("Part II:", lifeSupportRating(binaryNums[:]))
}
