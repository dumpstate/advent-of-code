import { getLines, toNumber } from "./common"

type Mapping = [string, string, number[][]]

function parse(lines: string[]): [number[], Mapping[]] {
	const seeds = lines[0].split(": ")[1].split(/\s+/).map(toNumber)
	const mappings: Mapping[] = []
	let mapping: Mapping = ["", "", []]

	for (let i = 2; i < lines.length; i++) {
		if (lines[i] === "") {
			mappings.push(mapping)
			mapping = ["", "", []]
			continue
		}

		if (lines[i].endsWith("map:")) {
			const split = lines[i].replace(" map:", "").split("-to-")
			mapping[0] = split[0]
			mapping[1] = split[1]
			continue
		}

		mapping[2].push(lines[i].split(/\s+/).map(toNumber))
	}

	return [seeds, mappings]
}

function find(
	mappings: Mapping[],
	from: string,
	to: string,
	value: number,
	reverse: boolean = false,
): number {
	let [curr, currVal] = [from, value]
	while (curr !== to) {
		for (const [f, t, rs] of mappings) {
			const [source, target] = reverse ? [t, f] : [f, t]
			if (curr !== source) {
				continue
			}

			for (let [d, s, span] of rs) {
				if (reverse) {
					;[s, d] = [d, s]
				}

				if (currVal >= s && currVal < s + span) {
					currVal = currVal - s + d
					break
				}
			}

			curr = target
			break
		}
	}
	return currVal
}

function isValidSeed(mappings: Mapping[], seedRs: number[], loc: number) {
	const seed = find(mappings, "location", "seed", loc, true)

	for (let i = 0; i < seedRs.length - 1; i += 2) {
		const [s, range] = [seedRs[i], seedRs[i + 1]]
		if (seed >= s && seed < s + range) {
			return true
		}
	}

	return false
}

function partI(input: [number[], Mapping[]]) {
	const [seeds, mappings] = input

	return Math.min.apply(
		null,
		seeds.map((seed) => find(mappings, "seed", "location", seed)),
	)
}

function partII(input: [number[], Mapping[]]) {
	const [seedRs, mappings] = input
	let loc = 0

	while (!isValidSeed(mappings, seedRs, loc)) {
		loc++
	}

	return loc
}

const mappings = parse(getLines(false))
console.log(`Part I: ${partI(mappings)}`)
console.log(`Part II: ${partII(mappings)}`)
