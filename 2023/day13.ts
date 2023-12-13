import { getLines, sum, transpose } from "./common"

function reflection(pattern: string[][], smudgeCount: number = 0): number {
	const smudges = {}
	for (let ix = 0; ix < pattern[0].length - 1; ix++) {
		smudges[ix] = 0
		for (let off = 0; off < pattern[0].length; off++) {
			const [i, j] = [ix - off, ix + off + 1]
			if (i < 0 || j >= pattern[0].length) {
				break
			}
			for (let k = 0; k < pattern.length; k++) {
				if (pattern[k][i] !== pattern[k][j]) {
					smudges[ix]++
				}
			}
		}
		if (smudges[ix] === smudgeCount) {
			return ix
		}
	}
	return -1
}

function score(pattern: string[][], smudgeCount: number = 0): number {
	const ver = reflection(pattern, smudgeCount)
	if (ver >= 0) {
		return ver + 1
	}
	return (reflection(transpose(pattern), smudgeCount) + 1) * 100
}

const patterns = getLines(false)
	.reduce(
		(acc: string[][], line: string) => {
			line === "" ? acc.push([]) : acc[acc.length - 1].push(line)
			return acc
		},
		[[]],
	)
	.filter((p) => p.length > 0)
	.map((p) => p.map((l) => l.split("")))
console.log(`Part I: ${sum(patterns.map((p) => score(p, 0)))}`)
console.log(`Part II: ${sum(patterns.map((p) => score(p, 1)))}`)
