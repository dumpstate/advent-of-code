package main

import (
	"fmt"
	"os"
	"strconv"

	"2021/common"
)

var BITS_CODE = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func toBinary(code string) string {
	binary := ""

	for _, char := range code {
		binary += BITS_CODE[char]
	}

	return binary
}

func toDecimal(binary string) int {
	value, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic("failed to convert to decimal")
	}

	return int(value)
}

func literalValue(binary string) (int, int) {
	ix := 0
	value := ""

	for {
		value += binary[ix+1 : ix+5]

		if binary[ix] == '0' {
			ix += 5
			break
		}

		ix += 5
	}

	return toDecimal(value), ix
}

func accFunc(typeId int) func(int, int) int {
	switch {
	case typeId == 0:
		return func(acc int, next int) int { return acc + next }
	case typeId == 1:
		return func(acc int, next int) int { return acc * next }
	case typeId == 2:
		return func(acc int, next int) int {
			if acc < next {
				return acc
			} else {
				return next
			}
		}
	case typeId == 3:
		return func(acc int, next int) int {
			if acc > next {
				return acc
			} else {
				return next
			}
		}
	case typeId == 5:
		return func(acc int, next int) int {
			if acc > next {
				return 1
			} else {
				return 0
			}
		}
	case typeId == 6:
		return func(acc int, next int) int {
			if acc < next {
				return 1
			} else {
				return 0
			}
		}
	case typeId == 7:
		return func(acc int, next int) int {
			if acc == next {
				return 1
			} else {
				return 0
			}
		}
	default:
		panic("unsupported packet type")
	}
}

func decode(binary string) (int, int, int) {
	version := toDecimal(binary[:3])
	packetType := toDecimal(binary[3:6])

	switch {
	case packetType == 4:
		value, nextOffset := literalValue(binary[6:])
		return version, value, 6 + nextOffset
	default:
		fn := accFunc(packetType)
		lengthTypeId := toDecimal(binary[6:7])
		switch {
		case lengthTypeId == 0:
			totalLength := 15
			subpacketOffset := 7 + totalLength
			subpacketLength := toDecimal(binary[7:subpacketOffset])
			versionTotal := version
			var acc int
			for ix := 0; true; ix++ {
				if subpacketOffset >= 7+totalLength+subpacketLength {
					break
				}

				subpacketVersion, value, nextOffset := decode(binary[subpacketOffset:])
				if ix == 0 {
					acc = value
				} else {
					acc = fn(acc, value)
				}

				versionTotal += subpacketVersion
				subpacketOffset += nextOffset
			}

			return versionTotal, acc, subpacketOffset
		case lengthTypeId == 1:
			totalLength := 11
			subpacketOffset := 7 + totalLength
			subpacketTotalCount := toDecimal(binary[7:subpacketOffset])
			versionTotal := version
			var acc int
			for ix := 0; ix < int(subpacketTotalCount); ix++ {
				subpacketVersion, value, nextOffset := decode(binary[subpacketOffset:])
				if ix == 0 {
					acc = value
				} else {
					acc = fn(acc, value)
				}

				subpacketOffset += nextOffset
				versionTotal += subpacketVersion
			}

			return versionTotal, acc, subpacketOffset
		default:
			panic("unsupported length type id")
		}
	}
}

func main() {
	code := common.ReadLines(os.Args[1])[0]
	version, value, _ := decode(toBinary(code))

	fmt.Println("Part I", version)
	fmt.Println("Part II", value)
}
