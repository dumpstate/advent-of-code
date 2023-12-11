import { getLines, getSize, l1 } from "./common"

function expansionIx(map: string[][]): [number[], number[]] {
	const [width, height] = getSize(map)
	const [cols, rows] = [[], []]

	for (let y = 0; y < height; y++) {
		if (map[y].every((c) => c === ".")) {
			rows.push(y)
		}
	}

	for (let x = 0; x < width; x++) {
		if (map.every((r) => r[x] === ".")) {
			cols.push(x)
		}
	}

	return [cols, rows]
}

function findGalaxies(map: string[][]): [number, number][] {
	const [width, height] = getSize(map)
	const galaxies: [number, number][] = []

	for (let y = 0; y < height; y++) {
		for (let x = 0; x < width; x++) {
			if (map[y][x] === "#") {
				galaxies.push([x, y])
			}
		}
	}

	return galaxies
}

function countExp(min: number, max: number, exp: number[]): number {
	let count = 0
	for (let i = 0; i < exp.length; i++) {
		if (exp[i] >= min && exp[i] < max) {
			count++
		}
	}
	return count
}

function l1WithExpansion(
	a: [number, number],
	b: [number, number],
	expCols: number[],
	expRows: number[],
	expSize: number,
): number {
	const [[ax, ay], [bx, by]] = [a, b]
	const xExpCount = countExp(Math.min(ax, bx), Math.max(ax, bx), expCols)
	const yExpCount = countExp(Math.min(ay, by), Math.max(ay, by), expRows)
	return l1(a, b) + (xExpCount + yExpCount) * expSize
}

function accDistance(
	map: string[][],
	galaxies: [number, number][],
	expSize: number,
) {
	const [expCols, expRows] = expansionIx(map)
	let acc = 0
	for (let i = 0; i < galaxies.length; i++) {
		for (let j = i + 1; j < galaxies.length; j++) {
			acc += l1WithExpansion(
				galaxies[i],
				galaxies[j],
				expCols,
				expRows,
				expSize,
			)
		}
	}
	return acc
}

const map = getLines().map((l) => l.split(""))
console.log(`Part I: ${accDistance(map, findGalaxies(map), 1)}`)
console.log(`Part II: ${accDistance(map, findGalaxies(map), 999999)}`)
