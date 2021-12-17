package main

import (
	"fmt"
	"os"
	"strings"

	"2021/common"
)

type Graph map[string]map[string]bool

func NewGraph(path string) *Graph {
	lines := common.ReadLines(path)
	graph := make(Graph)

	for _, line := range lines {
		split := strings.Split(line, "-")
		if len(split) != 2 {
			panic("invalid input")
		}

		from, to := split[0], split[1]
		_, existsFrom := graph[from]
		if !existsFrom {
			graph[from] = make(map[string]bool)
		}

		graph[from][to] = true

		_, existsTo := graph[to]
		if !existsTo {
			graph[to] = make(map[string]bool)
		}

		graph[to][from] = true
	}

	return &graph
}

func isComplete(path []string, endNode string) bool {
	return path[len(path)-1] == endNode
}

func incompletePaths(paths [][]string, endNode string) ([][]string, [][]string) {
	complete := make([][]string, 0)
	incomplete := make([][]string, 0)

	for _, path := range paths {
		if isComplete(path[:], endNode) {
			complete = append(complete, path)
		} else {
			incomplete = append(incomplete, path)
		}
	}

	return complete, incomplete
}

func isValid(path []string, nextNode string) bool {
	if common.IsLowerCased(nextNode) {
		return !common.ContainsStr(path, nextNode)
	}

	return true
}

func extendedIsValid(path []string, nextNode string) bool {
	switch {
	case nextNode == "start":
		return false
	case common.IsLowerCased(nextNode):
		count := common.MaxCountStr(
			common.FilterStr(path[:], common.IsLowerCased),
		)

		if count > 1 {
			return !common.ContainsStr(path, nextNode)
		} else {
			return true
		}
	default:
		return true
	}
}

func countAllPaths(
	graph *Graph,
	startNode string,
	endNode string,
	validator func([]string, string) bool,
) int {
	paths := [][]string{{startNode}}

	for {
		nextPaths, incomplete := incompletePaths(paths[:], endNode)
		if len(incomplete) == 0 {
			break
		}

		for _, path := range incomplete {
			for neighbour := range (*graph)[path[len(path)-1]] {
				if !validator(path[:], neighbour) {
					continue
				}

				nextPath := make([]string, len(path))
				copy(nextPath, path)

				nextPath = append(nextPath, neighbour)
				nextPaths = append(nextPaths, nextPath)
			}
		}

		paths = nextPaths
	}

	return len(paths)
}

func main() {
	graph := NewGraph(os.Args[1])

	fmt.Println("Part I", countAllPaths(graph, "start", "end", isValid))
	fmt.Println("Part II", countAllPaths(graph, "start", "end", extendedIsValid))
}
