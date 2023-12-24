import { Vec3D, getLines, toNumber } from "./common"

function intersects(
	[p1, v1]: [Vec3D, Vec3D],
	[p2, v2]: [Vec3D, Vec3D],
	min: number,
	max: number,
): boolean {
	const a = v1.y / v1.x
	const denom = a * v2.x - v2.y
	if (denom === 0) return false

	const t2 = (p2.y - p1.y + a * p1.x - a * p2.x) / denom
	if (t2 < 0) return false

	const t1 = (p2.x + v2.x * t2 - p1.x) / v1.x
	if (t1 < 0) return false

	const dx = p1.x + v1.x * t1
	const dy = p1.y + v1.y * t1

	return dx >= min && dx <= max && dy >= min && dy <= max
}

function partI(input: [Vec3D, Vec3D][], min: number, max: number) {
	let count = 0
	for (let i = 0; i < input.length; i++) {
		for (let j = i + 1; j < input.length; j++) {
			if (intersects(input[i], input[j], min, max)) count += 1
		}
	}
	return count
}

const input: [Vec3D, Vec3D][] = getLines()
	.map((l) => l.split(" @ "))
	.map(
		(s) =>
			s.map((e) => e.split(", ").map(toNumber)) as [
				[number, number, number],
				[number, number, number],
			],
	)
	.map(([p, v]) => [new Vec3D(...p), new Vec3D(...v)])
console.log("Part I", partI(input, 200000000000000, 400000000000000))
