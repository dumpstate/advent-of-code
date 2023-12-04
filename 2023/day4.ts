import { getLines, intersection, sum, toNumber } from "./common"

type Scratchcard = [number, number[], number[]]

function parse(line: string): Scratchcard {
	const [lStr, rStr] = line.split(":")
	const [winning, picked] = rStr.split(" | ")
	return [
		parseInt(lStr.substring(5).trim()),
		winning.trim().split(/\s+/).map(toNumber),
		picked.trim().split(/\s+/).map(toNumber),
	]
}

function score(s: Scratchcard): number {
	return intersection(s[1], s[2]).size
}

function partI(scratchcards: Scratchcard[]): number {
	return sum(
		scratchcards
			.map(score)
			.map((len) => (len === 0 ? 0 : Math.pow(2, len - 1))),
	)
}

function partII(scratchcards: Scratchcard[]): number {
	const cards: Map<number, number> = new Map()
	for (const s of scratchcards) {
		cards.set(s[0], 1)
	}

	for (let ix = 0; ix < scratchcards.length; ix++) {
		const s = scratchcards[ix]
		const count = cards.get(s[0]) as number
		for (let iy = s[0]; iy < s[0] + score(s); iy++) {
			const won = scratchcards[iy]
			cards.set(won[0], (cards.get(won[0]) as number) + count)
		}
	}

	return sum(Array.from(cards.values()))
}

const scratchcards = getLines().map(parse)
console.log(`Part I: ${partI(scratchcards)}`)
console.log(`Part II: ${partII(scratchcards)}`)
