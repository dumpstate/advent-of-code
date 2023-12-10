import fs from "fs"

export function getLines(filterEmpty: boolean = true): string[] {
	const lines = fs
		.readFileSync(process.argv[2] as string)
		.toString()
		.split("\n")

	if (filterEmpty) {
		return lines.filter((line) => line.length > 0)
	}

	return lines
}

export function sum(ns: number[]): number {
	return ns.reduce((acc, val) => acc + val, 0)
}

export function product(ns: number[]): number {
	return ns.reduce((acc, val) => acc * val, 1)
}

export function isDigit(c: string): boolean {
	return c >= "0" && c <= "9"
}

export function toNumber(c: string): number {
	return parseInt(c)
}

export function first<T>(ts: T[]): T {
	if (ts.length === 0) {
		throw new Error("Empty array")
	}

	return ts[0]
}

export function last<T>(ts: T[]): T {
	if (ts.length === 0) {
		throw new Error("Empty array")
	}

	return ts[ts.length - 1]
}

export function adjacentIx<T>(board: T[][], x: number, y: number): number[][] {
	return [
		[x - 1, y],
		[x + 1, y],
		[x, y - 1],
		[x, y + 1],
		[x - 1, y - 1],
		[x + 1, y - 1],
		[x - 1, y + 1],
		[x + 1, y + 1],
	].filter(
		([x, y]) => x >= 0 && x < board[0].length && y >= 0 && y < board.length,
	)
}

export function intersection<T>(l: T[], r: T[]): Set<T> {
	const lSet = new Set(l)
	return new Set(r.filter((val) => lSet.has(val)))
}

export function intersectAll<T>(...arrays: T[][]): Set<T> {
	return arrays.reduce(
		(acc, arr) => intersection(Array.from(acc), arr),
		new Set<T>(arrays[0]),
	)
}

export function zip<T, U>(l: T[], r: U[]): [T, U][] {
	return l.map((val, i) => [val, r[i]])
}

export function counter<T>(items: T[]): Map<T, number> {
	const counts = new Map<T, number>()
	for (const item of items) {
		counts.set(item, (counts.get(item) || 0) + 1)
	}
	return counts
}

export function keyOfMaxValue<T>(m: Map<string, T>): string {
	const keys = m.keys()
	let maxKey = keys.next().value
	for (const key of keys) {
		if ((m.get(key) as any) > (m.get(maxKey) as any)) {
			maxKey = key
		}
	}
	return maxKey
}

export function gcd(a: number, b: number): number {
	if (a < b) {
		;[a, b] = [b, a]
	}
	const r = a % b
	return r ? gcd(b, r) : b
}

export function lcm(a: number, b: number): number {
	return (a / gcd(a, b)) * b
}

export function lcmAll(ns: number[]): number {
	return ns.reduce(lcm)
}

export function deepClone<T>(v: T): T {
	return JSON.parse(JSON.stringify(v))
}
