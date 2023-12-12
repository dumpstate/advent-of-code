import { getLines, sum, toNumber } from "./common"

const cache = new Map<string, number>()

function count(hashLen: number, spring: string[], lens: number[]): number {
	const key = `${hashLen}:${spring.join(",")}:${lens.join(",")}`
	let res: number
	if (cache.has(key)) {
		return cache.get(key)
	}

	if (spring.length === 0) {
		if (lens.length === 1 && hashLen === lens[0]) {
			cache.set(key, 1)
			return 1
		}

		res = hashLen === 0 && lens.length === 0 ? 1 : 0
		cache.set(key, res)
		return res
	}

	switch (spring[0]) {
		case "?":
			res =
				count(hashLen, ["."].concat(spring.slice(1)), lens) +
				count(hashLen + 1, spring.slice(1), lens)
			break
		case ".":
			if (hashLen === 0) {
				res = count(0, spring.slice(1), lens)
			} else {
				res =
					hashLen === lens[0]
						? count(0, spring.slice(1), lens.slice(1))
						: 0
			}
			break
		case "#":
			res = count(hashLen + 1, spring.slice(1), lens)
			break
		default:
			throw new Error(`Unexpected character ${spring[0]}`)
	}
	cache.set(key, res)
	return res
}

function unfolded([spring, lens]: [string[], number[]]): [string[], number[]] {
	return [
		Array(5).fill(spring.join("")).join("?").split(""),
		Array(5).fill(lens.join(",")).join(",").split(",").map(toNumber),
	]
}

const input = getLines()
	.map((l) => l.split(" "))
	.map(
		([ss, ls]) =>
			[ss.split(""), ls.split(",").map(toNumber)] as [string[], number[]],
	)
console.log(`Part I: ${sum(input.map(([s, l]) => count(0, s, l)))}`)
console.log(
	`Part II: ${sum(input.map(unfolded).map(([s, l]) => count(0, s, l)))}`,
)
