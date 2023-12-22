import { deepClone, getLines, sum, toNumber } from "./common"

type Coord = [number, number, number]

function emptySpace(input: [Coord, Coord][]): number[][][] {
	let [maxX, maxY, maxZ] = [0, 0, 0]
	for (const [f, t] of input) {
		maxX = Math.max(maxX, f[0], t[0])
		maxY = Math.max(maxY, f[1], t[1])
		maxZ = Math.max(maxZ, f[2], t[2])
	}
	return Array(maxZ + 2)
		.fill(NaN)
		.map(() =>
			Array(maxY + 1)
				.fill(NaN)
				.map(() => Array(maxX + 1).fill(NaN)),
		)
}

function canFall(bid: number, [f, t], s: number[][][]) {
	const [[fx, fy, fz], [tx, ty, tz]] = [f, t]
	for (let z = fz; z <= tz; z++) {
		for (let y = fy; y <= ty; y++) {
			for (let x = fx; x <= tx; x++) {
				if (z <= 1) return false
				if (!isNaN(s[z - 1][y][x]) && s[z - 1][y][x] !== bid)
					return false
			}
		}
	}
	return true
}

function fall(input: [Coord, Coord][], s: number[][][]) {
	let tryFall = true
	while (tryFall) {
		let didMove = false
		for (let bid = 0; bid < input.length; bid++) {
			let [f, t] = input[bid]
			while (canFall(bid, [f, t], s)) {
				for (let z = f[2]; z <= t[2]; z++) {
					for (let y = f[1]; y <= t[1]; y++) {
						for (let x = f[0]; x <= t[0]; x++) {
							s[z][y][x] = NaN
							s[z - 1][y][x] = bid
						}
					}
				}
				f[2] -= 1
				t[2] -= 1
				didMove = true
			}
		}
		tryFall = didMove
	}
}

function countDisintegrated(
	bid: number,
	supportMapping: Record<number, Set<number>>,
) {
	const [disintegrated, visited] = [new Set(), new Set()]
	const q = [bid]
	while (q.length > 0) {
		const bid = q.shift()
		visited.add(bid)
		if (!supportMapping.hasOwnProperty(bid)) continue
		supportMapping[bid].forEach((did) => {
			disintegrated.add(did)
			if (!visited.has(did)) {
				q.push(did)
			}
		})
	}
	return disintegrated.size
}

function getSupportMapping(
	s: number[][][],
): [Record<number, Set<number>>, Set<number>] {
	const [canBeDisintegrated, supportMapping] = [new Set<number>(), {}]
	for (let z = 0; z < s.length; z++) {
		const bricks = new Set<number>()
		for (let y = 0; y < s[z].length; y++) {
			for (let x = 0; x < s[z][y].length; x++) {
				if (!isNaN(s[z][y][x])) {
					bricks.add(s[z][y][x])
				}
			}
		}

		for (const bid of bricks) {
			const [[fx, fy, _], [tx, ty, __]] = input[bid]
			for (let y = fy; y <= ty; y++) {
				for (let x = fx; x <= tx; x++) {
					const tid = s[z + 1][y][x]
					if (!isNaN(tid) && tid !== bid) {
						if (!supportMapping.hasOwnProperty(bid))
							supportMapping[bid] = new Set()
						supportMapping[bid].add(tid)
					}
				}
			}
			if (Object.keys(supportMapping).length === 0)
				canBeDisintegrated.add(bid)
			else canBeDisintegrated.delete(bid)
		}
	}
	for (const [bid, tids] of Object.entries(supportMapping)) {
		const tids_ = new Set()
		for (const tid of tids as Set<number>) {
			let found = false
			for (const [bid2, tids2] of Object.entries(supportMapping)) {
				if (bid2 === bid) continue
				if ((tids2 as Set<number>).has(tid)) {
					found = true
					break
				}
			}
			if (found) tids_.add(tid)
		}
		if (tids_.size === (tids as Set<number>).size)
			canBeDisintegrated.add(toNumber(bid))
	}
	for (let i = 0; i < input.length; i++)
		if (!supportMapping.hasOwnProperty(i)) canBeDisintegrated.add(i)

	return [supportMapping, canBeDisintegrated]
}

function plot(input: [Coord, Coord][], s: number[][][]) {
	for (let i = 0; i < input.length; i++) {
		const [[fx, fy, fz], [tx, ty, tz]] = input[i]
		for (let z = fz; z <= tz; z++) {
			for (let y = fy; y <= ty; y++) {
				for (let x = fx; x <= tx; x++) {
					s[z][y][x] = i
				}
			}
		}
	}
}

function partI(input: [Coord, Coord][]) {
	const s = emptySpace(input)
	plot(input, s)
	fall(deepClone(input), s)
	const [_, canBeDisintegrated] = getSupportMapping(s)
	return canBeDisintegrated.size
}

function partII(input: [Coord, Coord][]) {
	const s = emptySpace(input)
	plot(input, s)
	fall(deepClone(input), s)
	const [supportMapping, canBeDisintegrated] = getSupportMapping(s)

	return sum(
		input
			.map((_, i) => i)
			.filter((i) => !canBeDisintegrated.has(i))
			.map((i) => countDisintegrated(i, supportMapping)),
	)
}

const input = getLines().map(
	(l) =>
		l.split("~").map((c) => c.split(",").map(toNumber)) as [Coord, Coord],
)
console.log("Part I", partI(input))
console.log("Part II", partII(input))
