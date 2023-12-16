import { getLines } from "./common"

function toKey(pos: [number, number]): string {
	return `${pos[0]},${pos[1]}`
}

function advance(dir: string, [x, y]: [number, number]): [number, number] {
	switch (dir) {
		case ">":
			x++
			break
		case "<":
			x--
			break
		case "^":
			y--
			break
		case "v":
			y++
			break
	}
	return [x, y]
}

function mirror(tile: string, dir: string): string {
	switch (dir) {
		case ">":
			return tile === "/" ? "^" : "v"
		case "<":
			return tile === "/" ? "v" : "^"
		case "^":
			return tile === "/" ? ">" : "<"
		case "v":
			return tile === "/" ? "<" : ">"
		default:
			throw new Error(`Unknown direction ${dir}`)
	}
}

function nextDirs(tile: string, dir: string): string[] {
	if (tile === "/" || tile === "\\") {
		return [mirror(tile, dir)]
	}

	if (tile === "-") {
		if (dir === "^" || dir === "v") {
			return ["<", ">"]
		}
	}

	if (tile === "|") {
		if (dir === "<" || dir === ">") {
			return ["^", "v"]
		}
	}

	return [dir]
}

function simulate(
	cave: string[][],
	[initDir, [iX, iY]]: [string, [number, number]],
): number {
	const visited = new Set<string>()
	const beams: [string, [number, number]][] = nextDirs(
		cave[iY][iX],
		initDir,
	).map((d) => [d, [iX, iY]])
	visited.add(toKey([iX, iY]))

	while (beams.length > 0) {
		const [dir, [x, y]] = beams.shift()
		const [nX, nY] = advance(dir, [x, y])
		const key = toKey([nX, nY])
		if (nX < 0 || nX >= cave[0].length || nY < 0 || nY >= cave.length)
			continue

		if (
			((cave[nY][nX] === "|" && (dir === "<" || dir === ">")) ||
				(cave[nY][nX] === "-" && (dir === "^" || dir === "v"))) &&
			visited.has(key)
		)
			continue

		nextDirs(cave[nY][nX], dir).forEach((d) => beams.push([d, [nX, nY]]))
		visited.add(key)
	}

	return visited.size
}

function partI(cave: string[][]): number {
	return simulate(cave, [">", [0, 0]])
}

function partII(cave: string[][]): number {
	let max = 0
	for (let y = 0; y < cave.length; y++) {
		const l = simulate(cave, [">", [0, y]])
		const r = simulate(cave, ["<", [cave[y].length - 1, y]])
		max = Math.max(max, l, r)
	}
	for (let x = 0; x < cave[0].length; x++) {
		const t = simulate(cave, ["v", [x, 0]])
		const b = simulate(cave, ["^", [x, cave.length - 1]])
		max = Math.max(max, t, b)
	}
	return max
}

const input = getLines().map((l) => l.split(""))
console.log(`Part I: ${partI(input)}`)
console.log(`Part II: ${partII(input)}`)
