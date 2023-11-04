import fs from "fs"

export function getLines(): string[] {
	return fs
		.readFileSync(process.argv[2] as string)
		.toString()
		.split("\n")
		.filter((line) => line.length > 0)
}
