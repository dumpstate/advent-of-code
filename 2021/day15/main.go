package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"2021/common"
)

// https://pkg.go.dev/container/heap#example-package-PriorityQueue
type Item struct {
	value    common.Coord
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// increasing order
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value common.Coord, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func shortestPath(graph common.Matrix, source common.Coord, target common.Coord) []common.Coord {
	sizeX, sizeY := graph.Dimensions()
	dist, prev := make(map[common.Coord]int), make(map[common.Coord]common.Coord)
	visited := make(map[common.Coord]bool)
	vertexes := make(map[common.Coord]*Item)
	q := make(PriorityQueue, graph.TotalEntries())

	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			node := [2]int{x, y}
			dist[node] = math.MaxInt64

			vertexes[node] = &Item{
				value:    [2]int{x, y},
				priority: math.MaxInt64,
				index:    y*sizeX + x,
			}
			q[y*sizeX+x] = vertexes[node]
		}
	}
	q[0] = &Item{
		value:    source,
		priority: graph.GetOrPanic(source),
		index:    0,
	}
	dist[source] = graph.GetOrPanic(source)
	heap.Init(&q)

	for q.Len() > 0 {
		u := heap.Pop(&q).(*Item).value
		visited[u] = true

		if u == target {
			break
		}

		for _, neighbor := range graph.Neighbors(u) {
			_, exists := visited[neighbor]
			if exists {
				continue
			}

			alt := dist[u] + graph.GetOrPanic(neighbor)

			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = u
				q.update(vertexes[neighbor], neighbor, alt)
			}
		}
	}

	path, u := make([]common.Coord, 0), target

	for u != source {
		path = append(path, u)
		u = prev[u]
	}

	path = append(path, source)

	return path
}

func totalRisk(graph common.Matrix, path []common.Coord) int {
	risk := 0

	for _, node := range path[:len(path)-1] {
		risk += graph.GetOrPanic(node)
	}

	return risk
}

func resized(graph common.Matrix, times int) common.Matrix {
	sizeX, sizeY := graph.Dimensions()
	targetSizeX, targetSizeY := sizeX*times, sizeY*times
	target := make([][]int, targetSizeY)

	for y := 0; y < sizeY; y++ {
		for ty := 0; ty < times; ty++ {
			targetY := y + ty*sizeY
			target[targetY] = make([]int, targetSizeX)

			for x := 0; x < sizeX; x++ {
				for tx := 0; tx < times; tx++ {
					targetX := x + tx*sizeX
					target[targetY][targetX] = graph[y][x] + ty + tx
					if target[targetY][targetX] > 9 {
						target[targetY][targetX] = target[targetY][targetX]%10 + 1
					}
				}
			}
		}
	}

	return target
}

func main() {
	riskLevels := common.NewMatrix(os.Args[1], "")

	sizeX, sizeY := riskLevels.Dimensions()
	source, sink := [2]int{0, 0}, [2]int{sizeX - 1, sizeY - 1}
	path := shortestPath(riskLevels[:], source, sink)
	fmt.Println("Part I", totalRisk(riskLevels[:], path[:]))

	cave := resized(riskLevels[:], 5)
	caveSizeX, caveSizeY := cave.Dimensions()
	finalSink := [2]int{caveSizeX - 1, caveSizeY - 1}
	finalPath := shortestPath(cave[:], source, finalSink)
	fmt.Println("Part II", totalRisk(cave[:], finalPath[:]))
}
