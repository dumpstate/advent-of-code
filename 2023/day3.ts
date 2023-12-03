import { adjacentIx, getLines, isDigit, product, sum } from "./common"

function isSymbol(c: string): boolean {
	return !isDigit(c) && c !== "."
}

function checkHasAdjacentSymbol(
	board: string[][],
	x: number,
	y: number,
): boolean {
	return adjacentIx(board, x, y).some(([x, y]) => isSymbol(board[y][x]))
}

function parseNumbers(
	board: string[][],
	y: number,
	checkAdjacent: (b: string[][], x: number, y: number) => boolean,
): number[] {
	const res: number[] = []
	let currDigit: number[] = []
	let hasAdjacentSymbol = false

	function tryCollect() {
		if (currDigit.length > 0) {
			if (hasAdjacentSymbol) {
				res.push(parseInt(currDigit.join("")))
			}

			currDigit = []
			hasAdjacentSymbol = false
		}
	}

	for (let x = 0; x < board[y].length; x++) {
		const next = board[y][x]
		if (isDigit(next)) {
			currDigit.push(parseInt(next))

			if (!hasAdjacentSymbol && checkAdjacent(board, x, y)) {
				hasAdjacentSymbol = true
			}
		} else {
			tryCollect()
		}
	}

	tryCollect()

	return res
}

function findAdjacentNumbers(
	board: string[][],
	x: number,
	y: number,
): number[] {
	function isAdjacentWith(_: string[][], x_: number, y_: number): boolean {
		return adjacentIx(board, x_, y_).some(
			([x__, y__]) => x === x__ && y === y__,
		)
	}

	return [y - 1, y, y + 1]
		.filter((ix) => ix >= 0 && ix < board.length)
		.map((ix) => parseNumbers(board, ix, isAdjacentWith))
		.flat()
}

function partI(board: string[][]): number {
	const res: number[] = []

	for (let y = 0; y < board.length; y++) {
		res.push(...parseNumbers(board, y, checkHasAdjacentSymbol))
	}

	return sum(res)
}

function partII(board: string[][]): number {
	const ratios: number[] = []

	for (let y = 0; y < board.length; y++) {
		for (let x = 0; x < board[y].length; x++) {
			const next = board[y][x]
			if (next !== "*") {
				continue
			}

			const adjacentNumbers = findAdjacentNumbers(board, x, y)
			if (adjacentNumbers.length === 2) {
				ratios.push(product(adjacentNumbers))
			}
		}
	}

	return sum(ratios)
}

const board = getLines().map((line) => line.split(""))
console.log(`Part I: ${partI(board)}`)
console.log(`Part II: ${partII(board)}`)
