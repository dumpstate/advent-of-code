import { getLines, toNumber } from "./common"

type Direction = ">" | "<" | "^" | "v"
type Pos = [number, number]

function neighbors(
	board: number[][],
	dir: Direction,
	dirLen: number,
	[x, y]: Pos,
): [Direction, number, Pos][] {
	const ns = []

	if (dir === ">" || dir === "<") {
		if (y > 0) {
			ns.push(["^", 1, [x, y - 1]])
		}
		if (y < board.length - 1) {
			ns.push(["v", 1, [x, y + 1]])
		}
		if (dirLen < 3) {
			if (dir === "<" && x > 0) {
				ns.push([dir, dirLen + 1, [x - 1, y]])
			} else if (dir === ">" && x < board[0].length - 1) {
				ns.push([dir, dirLen + 1, [x + 1, y]])
			}
		}
	} else if (dir === "^" || dir === "v") {
		if (x > 0) {
			ns.push(["<", 1, [x - 1, y]])
		}
		if (x < board[0].length - 1) {
			ns.push([">", 1, [x + 1, y]])
		}
		if (dirLen < 3) {
			if (dir === "^" && y > 0) {
				ns.push([dir, dirLen + 1, [x, y - 1]])
			} else if (dir === "v" && y < board.length - 1) {
				ns.push([dir, dirLen + 1, [x, y + 1]])
			}
		}
	}

	return ns
}

function findMinPath(board: number[][]): number {
	const distances = {}
	for (const dir of [">", "<", "^", "v"]) {
		for (let i = 1; i <= 3; i++) {
			const init = Array(board.length)
			for (let y = 0; y < board.length; y++) {
				init[y] = Array(board[0].length).fill(Number.NaN)
			}
			distances[`${dir};${i}`] = init
		}
	}
	Object.values(distances).forEach((d) => (d[0][0] = 0))

	const q: [Direction, number, Pos][] = [[">", 1, [0, 0]]]

	while (q.length > 0) {
		const [dir, dirLen, [x, y]] = q.shift()
		const distKey = `${dir};${dirLen}`
		const dist = distances[distKey][y][x]

		for (const [nDir, nDirLen, [nX, nY]] of neighbors(board, dir, dirLen, [
			x,
			y,
		])) {
			const nKey = `${nDir};${nDirLen}`
			const nDist = distances[nKey][nY][nX]
			if (isNaN(nDist) || nDist > dist + board[nY][nX]) {
				distances[nKey][nY][nX] = dist + board[nY][nX]
				q.push([nDir, nDirLen, [nX, nY]])
			}
		}
	}

	return Math.min(
		...Object.values(distances)
			.map((d) => d[board.length - 1][board[0].length - 1])
			.filter((n) => !isNaN(n)),
	)
}

function partI(board: number[][]): number {
	return findMinPath(board)
}

const board = getLines().map((l) => l.split("").map(toNumber))
console.log(`Part I: ${partI(board)}`)
