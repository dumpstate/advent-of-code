import { first, getLines, last, sum, toNumber } from "./common"

function expand(ns: number[]): number[][] {
	const diffs: number[][] = [ns]
	while (last(diffs).some((n) => n !== 0)) {
		const next: number[] = []
		for (let i = 0; i < last(diffs).length - 1; i++) {
			next.push(last(diffs)[i + 1] - last(diffs)[i])
		}
		diffs.push(next)
	}
	return diffs.reverse()
}

const findNext = (ns: number[]) =>
	expand(ns).reduce((rem, diff) => last(diff) + rem, 0)
const findPrev = (ns: number[]) =>
	expand(ns).reduce((rem, diff) => first(diff) - rem, 0)

const input = getLines().map((l) => l.split(/\s+/).map(toNumber))
console.log(`Part I: ${sum(input.map(findNext))}`)
console.log(`Part II: ${sum(input.map(findPrev))}`)
