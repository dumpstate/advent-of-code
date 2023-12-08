import { getLines, lcmAll } from "./common"

type Network = { instr: string[]; map: Map<string, [string, string]> }

function countSteps(
	{ instr, map }: Network,
	start: string,
	end: (n: string) => boolean,
): number {
	let [i, curr] = [0, start]
	while (!end(curr)) {
		const s = instr[i % instr.length]
		const [l, r] = map.get(curr) as [string, string]
		curr = s === "L" ? l : r
		i += 1
	}
	return i
}

function partI(network: Network): number {
	return countSteps(network, "AAA", (n) => n === "ZZZ")
}

function partII(network: Network): number {
	return lcmAll(
		Array.from(network.map.keys())
			.filter((n) => n.endsWith("A"))
			.map((s) => countSteps(network, s, (n) => n.endsWith("Z"))),
	)
}

const lines = getLines()
const network = {
	instr: lines[0].split(""),
	map: lines
		.slice(1)
		.map((l) => [l.slice(0, 3), [l.slice(7, 10), l.slice(12, 15)]])
		.reduce((acc, [k, v]) => acc.set(k, v), new Map()),
}
console.log(`Part I: ${partI(network)}`)
console.log(`Part II: ${partII(network)}`)
