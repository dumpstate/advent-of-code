import { deepClone, getLines } from "./common"

function tilt(board: string[][]) {
	for (let j = 0; j < board.length; j++) {
		for (let i = 0; i < board[j].length; i++) {
			if (board[j][i] !== "O") {
				continue
			}

			for (let k = j - 1; k >= 0; k--) {
				if (board[k][i] !== ".") {
					break
				}

				board[k][i] = "O"
				board[k + 1][i] = "."
			}
		}
	}
}

function rotate(board: string[][]): string[][] {
	const rotated = Array(board.length)
		.fill(null)
		.map(() => Array(board[0].length).fill(null))
	for (let j = 0; j < board.length; j++) {
		for (let i = 0; i < board[0].length; i++) {
			rotated[i][board.length - 1 - j] = board[j][i]
		}
	}
	return rotated
}

function score(board: string[][]): number {
	let score = 0
	for (let j = 0; j < board.length; j++) {
		const rowScore = board.length - j
		for (let i = 0; i < board[0].length; i++) {
			if (board[j][i] === "O") {
				score += rowScore
			}
		}
	}
	return score
}

function partI(board: string[][]): number {
	tilt(board)
	return score(board)
}

function partII(board: string[][]): number {
	const seen = new Map()
	let ix = 0
	while (true) {
		for (let i = 0; i < 4; i++) {
			tilt(board)
			board = rotate(board)
		}
		const key = board.map((row) => row.join("")).join("")
		ix++
		if (seen.has(key)) {
			if ((1000000000 - ix) % (ix - seen.get(key)) === 0) {
				return score(board)
			}
		}
		seen.set(key, ix)
	}
}

const input = getLines().map((l) => l.split(""))
console.log(`Part I: ${partI(deepClone(input))}`)
console.log(`Part II: ${partII(deepClone(input))}`)
