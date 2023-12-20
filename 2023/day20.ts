import { getLines, lcmAll, sum } from "./common"

function flipFlop(name: string, ...targets: string[]) {
	let state = false

	return {
		type: "flipFlop",
		name,
		targets,
		reset: () => {
			state = false
		},
		accept: ([_, signal]: [string, number]) => {
			if (signal <= 0) {
				state = !state
				return targets.map((t) => [t, state ? 1 : 0])
			}
			return []
		},
	}
}

function conjunction(name: string, ...targets: string[]) {
	const state: Map<string, number> = new Map()

	return {
		type: "conjunction",
		name,
		targets,
		reset: () => {
			for (const key of state.keys()) {
				state.set(key, 0)
			}
		},
		connect: (...modules: string[]) => {
			modules.forEach((m) => state.set(m, 0))
		},
		accept: ([from, signal]: [string, number]) => {
			state.set(from, signal)
			if (sum(Array.from(state.values())) == state.size)
				return targets.map((t) => [t, 0])
			return targets.map((t) => [t, 1])
		},
	}
}

function broadcaster(...targets: string[]) {
	return {
		type: "broadcaster",
		name: "broadcaster",
		targets,
		reset: () => {},
		accept: ([_, signal]: [string, number]) =>
			targets.map((t) => [t, signal]),
	}
}

function simulate(
	modules,
	find: [string, number, string] | null = null,
): [number, number, boolean] {
	const curr: [string, string, number][] = [["button", "broadcaster", 0]]
	let [count, highSignals] = [0, 0]

	while (curr.length > 0) {
		const [source, target, signal] = curr.shift()
		count += 1
		highSignals += signal

		if (
			find &&
			source === find[0] &&
			signal === find[1] &&
			target === find[2]
		)
			return [highSignals, count - highSignals, true]
		if (!modules.hasOwnProperty(target)) continue
		modules[target]
			.accept([source, signal])
			.forEach(([t, s]) => curr.push([target, t, s]))
	}

	return [highSignals, count - highSignals, false]
}

function partI(modules) {
	let [highSignals, lowSignals] = [0, 0]
	for (let i = 0; i < 1000; i++) {
		const [hs, ls, _] = simulate(modules)
		highSignals += hs
		lowSignals += ls
	}
	return highSignals * lowSignals
}

function minCount(modules, findSource: [string, number, string]) {
	Object.values(modules).forEach((m: any) => m.reset())
	let count = 0
	while (true) {
		count += 1
		const [_, __, found] = simulate(modules, findSource)
		if (found) break
	}
	return count
}

function partII(modules) {
	// form input inspection:
	//     rx gets a signal only from xm
	//     xm is a conjunction module
	//     xm receives signal from sv and ng (botj conj.)
	//         => xm sends low signal only after receiving
	//            high signals from both sv, ng, ft, jz
	//     all under assumption we expect cycles
	return lcmAll([
		minCount(modules, ["sv", 1, "xm"]),
		minCount(modules, ["ng", 1, "xm"]),
		minCount(modules, ["ft", 1, "xm"]),
		minCount(modules, ["jz", 1, "xm"]),
	])
}

const modules = getLines()
	.map((l) => l.split(" -> "))
	.map(([type, targets]) => {
		const ts = targets.split(", ")
		if (type[0] === "%") {
			return flipFlop(type.slice(1), ...ts)
		}
		if (type[0] === "&") {
			return conjunction(type.slice(1), ...ts)
		}
		if (type === "broadcaster") {
			return broadcaster(...ts)
		}
		throw new Error(`Unknown type ${type}`)
	})
	.reduce((acc, m) => ({ ...acc, [m.name]: m }), {})
const conjModules = Object.values(modules)
	.filter((m: any) => m.type === "conjunction")
	.map((m: any) => m.name)
for (const m of Object.values(modules) as any) {
	for (const conj of conjModules) {
		if (m.targets.includes(conj)) {
			modules[conj].connect(m.name)
		}
	}
}
console.log("Part I", partI(modules))
console.log("Part II", partII(modules))
