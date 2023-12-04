import fs from "fs"

export function getLines(): string[] {
	return fs
		.readFileSync(process.argv[2] as string)
		.toString()
		.split("\n")
		.filter((line) => line.length > 0)
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