import { deepClone, first, getLines, toNumber } from "./common"

function find(
	pipes: string[][],
	fn: (item: string) => boolean,
	start: number = 0,
	step: number = 1,
): [number, number][] {
	const matching: [number, number][] = []
	for (let y = start; y < pipes.length; y += step) {
		for (let x = start; x < pipes[y].length; x += step) {
			if (fn(pipes[y][x])) {
				matching.push([x, y])
			}
		}
	}
	return matching
}

const NEIGHBOR = {
	N: new Set(["|", "7", "F"]),
	S: new Set(["|", "L", "J"]),
	W: new Set(["-", "F", "L"]),
	E: new Set(["-", "J", "7"]),
}

function findDirs(pipes: string[][], [x, y]: [number, number]): string[] {
	return Object.entries(NEIGHBOR)
		.filter(([dir, items]) => {
			const [nX, nY] = coord([x, y], dir)
			if (nX < 0 || nY < 0 || nY >= pipes.length || nX >= pipes[0].length)
				return false
			return items.has(pipes[nY][nX])
		})
		.map(([dir, _]) => dir)
}

function coord([x, y]: [number, number], dir: string): [number, number] {
	switch (dir) {
		case "N":
			return [x, y - 1]
		case "S":
			return [x, y + 1]
		case "W":
			return [x - 1, y]
		case "E":
			return [x + 1, y]
		default:
			throw new Error(`Not a direction: ${dir}`)
	}
}

function toKey([x, y]: [number, number]): string {
	return `${x},${y}`
}

function fromKey(key: string): [number, number] {
	return key.split(",").map(toNumber) as [number, number]
}

function findPipeline(
	pipes: string[][],
	s: [number, number],
): Map<string, number> {
	const board = deepClone(pipes)
	const toVisit: [[number, number], number][] = [[s, 0]]
	const visited = new Map<string, number>()
	visited.set(toKey(s), 0)

	while (toVisit.length > 0) {
		const [[x, y], dist] = toVisit.shift() as any
		board[y][x] = dist
		visited.set(toKey([x, y]), dist)
		findDirs(board, [x, y]).forEach((n) => {
			toVisit.push([coord([x, y], n), dist + 1])
		})
	}

	return visited
}

function markInflated(
	pipes: string[][],
	[x, y]: [number, number],
	curr: string,
) {
	switch (curr) {
		case "|":
			pipes[y - 1][x] = curr
			pipes[y][x] = curr
			pipes[y + 1][x] = curr
			break
		case "-":
			pipes[y][x - 1] = curr
			pipes[y][x] = curr
			pipes[y][x + 1] = curr
			break
		case "L":
			pipes[y - 1][x] = "|"
			pipes[y][x] = curr
			pipes[y][x + 1] = "-"
			break
		case "J":
			pipes[y - 1][x] = "|"
			pipes[y][x] = curr
			pipes[y][x - 1] = "-"
			break
		case "7":
			pipes[y][x - 1] = "-"
			pipes[y][x] = curr
			pipes[y + 1][x] = "|"
			break
		case "F":
			pipes[y][x + 1] = "-"
			pipes[y][x] = curr
			pipes[y + 1][x] = "|"
			break
	}
}

function findType(pipes: string[][], [x, y]: [number, number]): string {
	const dirs = new Set<string>(findDirs(pipes, [x, y]))
	if (dirs.has("N") && dirs.has("S")) return "|"
	if (dirs.has("W") && dirs.has("E")) return "-"
	if (dirs.has("N") && dirs.has("E")) return "L"
	if (dirs.has("N") && dirs.has("W")) return "J"
	if (dirs.has("S") && dirs.has("W")) return "7"
	if (dirs.has("S") && dirs.has("E")) return "F"
	throw new Error("Cannot identify the type")
}

function inflate(pipes: string[][]): string[][] {
	const res: string[][] = Array(pipes.length * 3)
	for (let y = 0; y < res.length; y++) {
		res[y] = Array.from(".".repeat(pipes[0].length * 3))
	}

	for (let y = 0; y < pipes.length; y++) {
		for (let x = 0; x < pipes[y].length; x++) {
			let curr = pipes[y][x]
			if (curr === "S") {
				curr = findType(pipes, [x, y])
			}
			markInflated(res, [x * 3 + 1, y * 3 + 1], curr)
		}
	}

	return res
}

function flood(
	board: string[][],
	shoelace: Set<string>,
	fillChar: string = "O",
) {
	const toVisit: [number, number][] = [[0, 0]]
	const visited = new Set<string>()

	while (toVisit.length > 0) {
		const [x, y] = toVisit.shift()
		const key = toKey([x, y])
		if (
			shoelace.has(key) ||
			visited.has(key) ||
			x < 0 ||
			y < 0 ||
			y > board.length - 1 ||
			x > board[0].length - 1
		) {
			continue
		}
		board[y][x] = fillChar
		visited.add(key)
		toVisit.push([x - 1, y], [x + 1, y], [x, y - 1], [x, y + 1])
	}
}

function fill(board: string[][], keys: Set<string>, fillChar: string = "*") {
	for (const key of keys) {
		const [x, y] = fromKey(key)
		board[y][x] = fillChar
	}
}

function partI(pipes: string[][]): number {
	return Math.max(
		...findPipeline(pipes, first(find(pipes, (c) => c === "S"))).values(),
	)
}

function partII(pipes: string[][]): number {
	const [sx, sy] = first(find(pipes, (c) => c === "S"))
	const inflated = inflate(pipes)
	const pipeline = new Set(
		findPipeline(inflated, [sx * 3 + 1, sy * 3 + 1]).keys(),
	)

	flood(inflated, pipeline, "O")
	fill(inflated, pipeline, "*")

	return find(inflated, (c) => c !== "*" && c !== "O", 1, 3).length
}

const input = getLines().map((l) => l.split(""))
console.log(`Part I: ${partI(deepClone(input))}`)
console.log(`Part II: ${partII(deepClone(input))}`)
