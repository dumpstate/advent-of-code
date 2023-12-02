import { getLines, product, sum, toNumber } from "./common"

type Game = [number, number[][]]

function parse(line: string): Game {
	const [_, id, cs]: string[] = line.match(/Game (\d+): (.*)/) || []
	return [
		toNumber(id),
		cs
			.split("; ")
			.map((cl) =>
				[
					cl.match(/(\d+) red/)?.at(1) || "0",
					cl.match(/(\d+) blue/)?.at(1) || "0",
					cl.match(/(\d+) green/)?.at(1) || "0",
				].map(toNumber),
			),
	]
}

function partI(gs: Game[]): number {
	return sum(
		gs
			.filter(
				([_, cs]) =>
					cs.filter(([r, b, g]) => r <= 12 && b <= 14 && g <= 13)
						.length === cs.length,
			)
			.map(([id, _]) => id),
	)
}

function partII(gs: Game[]): number {
	return sum(
		gs.map(([_, cs]) =>
			product(
				cs.reduce(
					([rMax, bMax, gMax], [r, b, g]) => [
						Math.max(rMax, r),
						Math.max(bMax, b),
						Math.max(gMax, g),
					],
					cs[0],
				),
			),
		),
	)
}

const input = getLines().map(parse)
console.log(`Part I: ${partI(input)}`)
console.log(`Part II: ${partII(input)}`)
