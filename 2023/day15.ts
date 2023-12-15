import { getLines, product, sum } from "./common"

function hash(str: string): number {
	let res = 0
	for (let i = 0; i < str.length; i++) {
		res += str.charCodeAt(i)
		res *= 17
		res = res % 256
	}
	return res
}

function partI(input: string[]): number {
	return sum(input.map(hash))
}

function partII(input: string[]): number {
	const instructions: [string, string, number][] = input
		.map((s) => s.match(/(.*)(=|-)(\d*)/).slice(1, 4))
		.map((s) => [s[0], s[1], s[2] === "" ? 0 : parseInt(s[2])])
	const boxes = {}
	instrloop: for (const [lensLabel, op, val] of instructions) {
		const boxId = hash(lensLabel)
		if (!boxes.hasOwnProperty(boxId)) {
			boxes[boxId] = []
		}

		switch (op) {
			case "-":
				for (let i = 0; i < boxes[boxId].length; i++) {
					if (boxes[boxId][i][0] === lensLabel) {
						boxes[boxId] = boxes[boxId]
							.slice(0, i)
							.concat(boxes[boxId].slice(i + 1))
						break
					}
				}
				break
			case "=":
				for (let i = 0; i < boxes[boxId].length; i++) {
					if (boxes[boxId][i][0] === lensLabel) {
						boxes[boxId][i][1] = val
						continue instrloop
					}
				}

				boxes[boxId].push([lensLabel, val])
				break
		}
	}

	return sum(
		Object.entries(boxes)
			.flatMap(([boxId, lenses]: [string, [string, number][]]) =>
				lenses.map(([_, lens], ix) => [
					parseInt(boxId) + 1,
					lens,
					ix + 1,
				]),
			)
			.map(product),
	)
}

const input = getLines().flatMap((l) => l.split(","))
console.log(`Part I: ${partI(input)}`)
console.log(`Part II: ${partII(input)}`)
