import { getLines } from "./common"

function asInstr(col: string): [string, number] {
	return [
		{ "0": "R", "1": "D", "2": "L", "3": "U" }[col.slice(-1)[0]],
		parseInt(col.slice(1, -1), 16),
	]
}

function vertexes(instr: [string, number][]): [number, number][] {
	const vs: [number, number][] = [[0, 0]]
	let [x, y] = vs[0]
	for (const [d, l] of instr) {
		if (d === "R") x += l
		if (d === "L") x -= l
		if (d === "U") y -= l
		if (d === "D") y += l
		vs.push([x, y])
	}
	return vs
}

function area(instr: [string, number][]): number {
	const vs = vertexes(instr)
	let [a, b] = [0, 0]
	for (let i = 0; i < vs.length - 1; i++) {
		const [[x1, y1], [x2, y2]] = [vs[i], vs[i + 1]]
		// shoelace formula
		a += x1 * y2 - x2 * y1
		// boundary
		b += Math.abs(x2 - x1) + Math.abs(y2 - y1)
	}
	// from Pick's theorem
	return b + a / 2 - b / 2 + 1
}

const input: [string, number, string][] = getLines()
	.map((l) => l.split(" "))
	.map(([dir, len, col]) => [dir, Number(len), col.slice(1, -1)])
console.log(`Part I: ${area(input.map(([d, l, _]) => [d, l]))}`)
console.log(`Part II: ${area(input.map(([_, __, c]) => asInstr(c)))}`)
