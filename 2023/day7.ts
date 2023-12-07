import { counter, getLines, keyOfMaxValue, sum, zip } from "./common"

type Hand = [string[], number]

const STRENGTH = {
	A: 14,
	K: 13,
	Q: 12,
	J: 11,
	T: 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

function getType([h, _]: Hand): number {
	const c = counter(h)
	const vs = new Set(c.values())
	if (vs.has(5)) return 7 // five of a kind
	if (vs.has(4)) return 6 // four of a kind
	if (vs.has(3) && vs.has(2)) return 5 // full house
	if (vs.has(3)) return 4 // three of a kind
	if (vs.has(2) && vs.has(1) && c.size === 3) return 3 // two pair
	if (vs.has(2)) return 2 // one pair
	return 1
}

function getTypeJ([h, bid]: Hand): number {
	const maxC = keyOfMaxValue(counter(h.filter((v) => v !== "J")))
	return getType([h.map((v) => (v === "J" ? maxC : v)), bid])
}

function compareStrength(
	strength: { [key: string]: number },
	[l, _]: Hand,
	[r, __]: Hand,
) {
	for (let i = 0; i < l.length; i++) {
		if (l[i] === r[i]) continue
		return strength[l[i]] - strength[r[i]]
	}

	return 0
}

const compare =
	(typeGetter: (h: Hand) => number, strength: { [key: string]: number }) =>
	(l: Hand, r: Hand) => {
		const [lt, rt] = [typeGetter(l), typeGetter(r)]
		if (lt !== rt) return lt - rt
		return compareStrength(strength, l, r)
	}

const score = (hs: Hand[]) =>
	sum(hs.map(([_, b]: Hand, ix: number) => (ix + 1) * b))

function partI(hs: Hand[]): number {
	hs.sort(compare(getType, STRENGTH))
	return score(hs)
}

function partII(hs: Hand[]): number {
	hs.sort(compare(getTypeJ, { ...STRENGTH, J: 1 }))
	return score(hs)
}

const hands: Hand[] = getLines()
	.map((l) => l.split(" "))
	.map(([h, b]) => [h.split(""), parseInt(b)])
console.log(`Part I: ${partI([...hands])}`)
console.log(`Part II: ${partII([...hands])}`)
