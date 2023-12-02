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
