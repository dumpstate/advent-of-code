import { first, getLines, isDigit, last, sum, toNumber } from "./common"

const DIGITS = {
	one: 1,
	two: 2,
	three: 3,
	four: 4,
	five: 5,
	six: 6,
	seven: 7,
	eight: 8,
	nine: 9,
}

function calibrationValue(ns: number[]): number {
	return first(ns) * 10 + last(ns)
}

function findDigits(line: string): number[] {
	const digits: number[] = []
	for (let i = 0; i < line.length; i++) {
		if (isDigit(line[i])) {
			digits.push(toNumber(line[i]))
			continue
		}

		for (const [word, digit] of Object.entries(DIGITS)) {
			if (line.startsWith(word, i)) {
				digits.push(digit)
				break
			}
		}
	}
	return digits
}

function part1(lines: string[]): number {
	return sum(
		lines
			.map((line) => Array.from(line).filter(isDigit).map(toNumber))
			.map(calibrationValue),
	)
}

function part2(lines: string[]): number {
	return sum(lines.map(findDigits).map(calibrationValue))
}

const lines = getLines()
console.log(`Part I: ${part1(lines)}`)
console.log(`Part II: ${part2(lines)}`)
