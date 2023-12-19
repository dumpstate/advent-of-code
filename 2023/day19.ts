import { deepClone, getLines, product, sum } from "./common"

type Rating = { x: number; m: number; a: number; s: number }

function parseRating(line: string): Rating {
	return line
		.slice(1, -1)
		.split(",")
		.reduce((acc, r) => {
			const [k, v] = r.split("=")
			return { ...acc, [k]: parseInt(v) }
		}, {}) as Rating
}

function appendWorkflow(workflows: Object, line: string) {
	const [name, expr] = line.split("{")
	const ex = expr.slice(0, -1).split(",")
	const rules = ex.slice(0, -1).map((r) => {
		const [ex, t] = r.split(":")
		const op = ex.includes(">") ? ">" : "<"
		const [property, value] = ex.split(op)
		return {
			rules: [{ op, property, value: parseInt(value) }],
			target: t,
		}
	})

	workflows[name] = {
		rules,
		target: ex.slice(-1)[0],
	}
}

function evaluate(workflows: Object, rating: Rating, state: string) {
	if (state === "A" || state === "R") return state

	const { rules, target } = workflows[state]
	for (const rule of rules) {
		for (const { op, property, value } of rule.rules) {
			if (op === ">" && rating[property] > value)
				return evaluate(workflows, rating, rule.target)
			if (op === "<" && rating[property] < value)
				return evaluate(workflows, rating, rule.target)
		}
	}
	return evaluate(workflows, rating, target)
}

function compliment(rs) {
	return rs.rules.map((r) => ({
		...r,
		op: r.op === ">" ? "<" : ">",
		value: r.op === ">" ? r.value + 1 : r.value - 1,
	}))[0]
}

function count(rs) {
	return product(Object.values(rs).map(([f, t]) => Math.max(0, t - f + 1)))
}

function partI({
	workflows,
	ratings,
}: {
	workflows: Object
	ratings: Rating[]
}) {
	return sum(
		ratings
			.map((r) => ({ rating: r, result: evaluate(workflows, r, "in") }))
			.filter(({ result }) => result === "A")
			.flatMap(({ rating }) => Object.values(rating)),
	)
}

function applyRules(ranges, rules) {
	const rs = deepClone(ranges)
	for (const { op, property, value } of rules) {
		if (op === ">") rs[property][0] = Math.max(rs[property][0], value + 1)
		if (op === "<") rs[property][1] = Math.min(rs[property][1], value - 1)
	}
	return rs
}

function partII({ workflows }: { workflows: Object }) {
	const q: [string, any][] = [
		["in", { x: [1, 4000], m: [1, 4000], a: [1, 4000], s: [1, 4000] }],
	]
	let score = 0

	while (q.length > 0) {
		const [state, rs] = q.shift()
		if (
			Object.values(rs)
				.map(([f, t]) => t - f + 1)
				.some((n) => n <= 0)
		)
			continue

		const { rules, target } = workflows[state]
		for (const rule of rules) {
			if (rule.target === "A") {
				score += count(applyRules(rs, rule.rules))
			} else if (rule.target !== "R") {
				q.push([rule.target, applyRules(rs, rule.rules)])
			}
		}

		if (target === "A") {
			score += count(applyRules(rs, rules.map(compliment)))
		} else if (target !== "R") {
			q.push([target, applyRules(rs, rules.map(compliment))])
		}
	}

	return score
}

const input = getLines(false).reduce(
	({ workflows, ratings }, line) => {
		if (ratings === null) {
			if (line === "") return { workflows, ratings: [] }
			appendWorkflow(workflows, line)
		} else if (line != "") ratings.push(parseRating(line))
		return { workflows, ratings }
	},
	{ workflows: {}, ratings: null },
)
console.log(`Part I: ${partI(input)}`)
console.log(`Part II: ${partII(input)}`)
