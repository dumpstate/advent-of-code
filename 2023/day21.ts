import { getLines, getSize } from "./common"

function start(board: string[][]): [number, number] {
	for (let y = 0; y < board.length; y++) {
		for (let x = 0; x < board[y].length; x++) {
			if (board[y][x] === "S") {
				return [x, y]
			}
		}
	}
	throw new Error("S not found")
}

function countSteps(
	board: string[][],
	maxSteps: number,
	[sx, sy]: [number, number],
): number {
	const [width, height] = getSize(board)
	const dists = {}
	dists[`${sx},${sy}`] = 0
	const q: [number, number][] = [[sx, sy]]

	while (q.length > 0) {
		const [x, y] = q.shift()
		const dist = dists[`${x},${y}`]
		if (dist >= maxSteps) continue

		for (const [nx, ny] of [
			[x - 1, y],
			[x + 1, y],
			[x, y - 1],
			[x, y + 1],
		]) {
			const [bnx, bny] = [
				(nx + width * (Math.floor(Math.abs(nx / width)) + 1)) % width,
				(ny + height * (Math.floor(Math.abs(ny / height)) + 1)) %
					height,
			]
			if (board[bny][bnx] === "#") continue
			const key = `${nx},${ny}`
			if (!dists.hasOwnProperty(key) || dists[key] > dist + 1) {
				dists[key] = dist + 1
				q.push([nx, ny])
			}
		}
	}

	let count = 0
	for (const dist of Object.values(dists) as any) {
		if (
			dist == maxSteps ||
			(dist < maxSteps && (maxSteps - dist) % 2 === 0)
		) {
			count += 1
		}
	}
	return count
}

const input = getLines().map((l) => l.split(""))
console.log("Part I", countSteps(input, 64, start(input)))
console.log("Part II", countSteps(input, 26501365, start(input)))
