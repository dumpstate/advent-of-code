import { first, getLines, last } from "./common"

function find(row: string[], char: string) {
	return row.findIndex((c) => c === char)
}

function neighbours(input: string[][], [x, y]: [number, number]) {
	switch (input[y][x]) {
		case ".":
			return [
				[x - 1, y],
				[x + 1, y],
				[x, y - 1],
				[x, y + 1],
			]
		case ">":
			return [[x + 1, y]]
		case "<":
			return [[x - 1, y]]
		case "^":
			return [[x, y - 1]]
		case "v":
			return [[x, y + 1]]
		default:
			throw new Error("Unknown tile")
	}
}

function findAllPaths(input: string[][]) {
	const paths = []
	const [sX, eX] = [find(first(input), "."), find(last(input), ".")]
	const q: [[number, number], Set<string>][] = [[[sX, 0], new Set<string>()]]

	while (q.length > 0) {
		const [[x, y], path] = q.shift()
		if (y === input.length - 1 && x === eX) {
			paths.push(path)
			continue
		}
		path.add(`${x},${y}`)

		for (const [nx, ny] of neighbours(input, [x, y])) {
			if (
				nx < 0 ||
				nx >= input[0].length ||
				ny < 0 ||
				ny >= input.length ||
				input[ny][nx] === "#" ||
				path.has(`${nx},${ny}`)
			)
				continue

			q.push([[nx, ny], new Set(Array.from(path))])
		}
	}

	return paths
}

function partI(input: string[][]) {
	return Math.max(...findAllPaths(input).map((p) => p.size))
}

const input = getLines().map((l) => l.split(""))
console.log("Part I", partI(input))
