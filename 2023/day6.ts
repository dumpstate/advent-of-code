import { getLines, product, toNumber, zip } from "./common"

function combinations([t, d]: [number, number]): number {
	let count = 0

	for (let i = 0; i < t; i++) {
		if (i * (t - i) > d) {
			count += 1
		}
	}

	return count
}

function partI(lines: string[]): number {
	const [ts, ds] = lines.map((l) =>
		l.split(":")[1].trim().split(/\s+/).map(toNumber),
	)
	return product(zip(ts, ds).map(combinations))
}

function partII(lines: string[]): number {
	const [t, d] = lines.map((l) => parseInt(l.split(":")[1].replace(/ /g, "")))
	return combinations([t, d])
}

const lines = getLines()
console.log(`Part I: ${partI(lines)}`)
console.log(`Part II: ${partII(lines)}`)
