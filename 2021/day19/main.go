package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"2021/common"
)

type Coord struct {
	x, y, z int
}

func NewCoord(line string) *Coord {
	split := strings.Split(line, ",")

	return &Coord{
		x: common.StrToInt(split[0]),
		y: common.StrToInt(split[1]),
		z: common.StrToInt(split[2]),
	}
}

func parse(lines []string) map[int][]*Coord {
	scanners := make(map[int][]*Coord)
	scannerId := 0
	scanner := make([]*Coord, 0)

	for _, line := range lines {
		if line == "" {
			scanners[scannerId] = scanner
			scannerId++
			scanner = make([]*Coord, 0)
			continue
		}

		if line[:3] == "---" {
			continue
		}

		scanner = append(scanner, NewCoord(line))
	}

	scanners[scannerId] = scanner

	return scanners
}

var Orientations map[int]func(Coord) Coord = map[int]func(Coord) Coord{
	0:  func(c Coord) Coord { return Coord{x: c.x, y: c.y, z: c.z} },
	1:  func(c Coord) Coord { return Coord{x: -c.z, y: -c.y, z: -c.x} },
	2:  func(c Coord) Coord { return Coord{x: -c.x, y: c.y, z: -c.z} },
	3:  func(c Coord) Coord { return Coord{x: c.x, y: -c.y, z: -c.z} },
	4:  func(c Coord) Coord { return Coord{x: -c.x, y: c.z, z: c.y} },
	5:  func(c Coord) Coord { return Coord{x: c.x, y: c.z, z: -c.y} },
	6:  func(c Coord) Coord { return Coord{x: c.z, y: c.y, z: -c.x} },
	7:  func(c Coord) Coord { return Coord{x: c.y, y: c.z, z: c.x} },
	8:  func(c Coord) Coord { return Coord{x: c.y, y: -c.z, z: -c.x} },
	9:  func(c Coord) Coord { return Coord{x: c.y, y: c.x, z: -c.z} },
	10: func(c Coord) Coord { return Coord{x: c.x, y: -c.z, z: c.y} },
	11: func(c Coord) Coord { return Coord{x: -c.x, y: -c.z, z: -c.y} },
	12: func(c Coord) Coord { return Coord{x: -c.y, y: -c.z, z: c.x} },
	13: func(c Coord) Coord { return Coord{x: -c.y, y: -c.x, z: -c.z} },
	14: func(c Coord) Coord { return Coord{x: -c.x, y: -c.y, z: c.z} },
	15: func(c Coord) Coord { return Coord{x: -c.y, y: c.x, z: c.z} },
	16: func(c Coord) Coord { return Coord{x: c.z, y: c.x, z: c.y} },
	17: func(c Coord) Coord { return Coord{x: c.y, y: -c.x, z: c.z} },
	18: func(c Coord) Coord { return Coord{x: -c.z, y: c.x, z: -c.y} },
	19: func(c Coord) Coord { return Coord{x: -c.y, y: c.z, z: -c.x} },
	20: func(c Coord) Coord { return Coord{x: c.z, y: -c.y, z: c.x} },
	21: func(c Coord) Coord { return Coord{x: c.z, y: -c.x, z: -c.y} },
	22: func(c Coord) Coord { return Coord{x: -c.z, y: -c.x, z: c.y} },
	23: func(c Coord) Coord { return Coord{x: -c.z, y: c.y, z: c.x} },
}

func square(num int) float64 {
	return float64(num * num)
}

func euclidean(first *Coord, second *Coord) float64 {
	return math.Sqrt(
		square(first.x-second.x) +
			square(first.y-second.y) +
			square(first.z-second.z),
	)
}

func manhattan(first *Coord, second *Coord) int {
	return common.Abs(first.x, second.x) + common.Abs(first.y, second.y) + common.Abs(first.z, second.z)
}

func distanceMatrix(reads []*Coord) *[][]float64 {
	distances := make([][]float64, len(reads))

	for ix, first := range reads {
		distances[ix] = make([]float64, len(reads))

		for iy, second := range reads {
			distances[ix][iy] = euclidean(first, second)
		}
	}

	return &distances
}

func findOverlaps(first []float64, second []float64) map[float64]bool {
	overlaps := make(map[float64]bool)
	firstMap := make(map[float64]bool)

	for _, value := range first {
		firstMap[value] = true
	}

	for _, value := range second {
		if value == 0 {
			continue
		}

		_, exists := firstMap[value]
		if exists {
			overlaps[value] = true
		}
	}

	return overlaps
}

func findMaxOverlap(first *[][]float64, second *[][]float64) (map[float64]bool, int, int) {
	maxOverlap := math.MinInt64
	var overlap map[float64]bool
	var resA, resB int

	for a, firstDistances := range *first {
		for b, secondDistances := range *second {
			if a == b {
				continue
			}

			overlaps := findOverlaps(firstDistances[:], secondDistances[:])

			if len(overlaps) > maxOverlap {
				maxOverlap = len(overlaps)
				overlap = overlaps
				resA = a
				resB = b
			}
		}
	}

	return overlap, resA, resB
}

func findScannerPosition(
	scanners *map[int][]*Coord,
	distances *[]*[][]float64,
	first int,
	second int,
	firstScanner Scanner,
) Scanner {
	_, p1, p2 := findMaxOverlap((*distances)[first], (*distances)[second])
	pairs := make(map[Coord]Coord)

	for a, dist1 := range (*(*distances)[first])[p1] {
		for b, dist2 := range (*(*distances)[second])[p2] {
			if dist1 == dist2 {
				sourceCoord := *(*scanners)[first][a]
				targetCoord := *(*scanners)[second][b]

				pairs[sourceCoord] = targetCoord
			}
		}
	}

	for tfid, ofn := range Orientations {
		var candidate *Coord = nil
		matchingPairs := 0

		for s, t := range pairs {
			source := Orientations[firstScanner.orientation](s)
			target := ofn(t)

			if candidate == nil {
				candidate = &Coord{
					x: source.x - target.x,
					y: source.y - target.y,
					z: source.z - target.z,
				}

				matchingPairs++
			} else {
				check := Coord{
					x: source.x - candidate.x,
					y: source.y - candidate.y,
					z: source.z - candidate.z,
				}

				if check.x == target.x && check.y == target.y && check.z == target.z {
					matchingPairs++
				} else {
					break
				}
			}
		}

		if matchingPairs == len(pairs) {
			candidate.x += firstScanner.pos.x
			candidate.y += firstScanner.pos.y
			candidate.z += firstScanner.pos.z

			return Scanner{pos: *candidate, orientation: tfid}
		}

		candidate = nil
	}

	panic("position not found")
}

func getPlan(
	scanners *map[int][]*Coord,
	distances *[]*[][]float64,
) [][2]int {
	plan := make([][2]int, 0)
	known := map[int]bool{0: true}
	added := true

	for {
		added = false

		for scannerIx := range known {
			for targetIx := 0; targetIx < len(*scanners); targetIx++ {
				overlaps, _, _ := findMaxOverlap((*distances)[scannerIx], (*distances)[targetIx])
				if len(overlaps) < 11 {
					continue
				}

				_, existsX := known[scannerIx]
				_, existsY := known[targetIx]

				if existsX && !existsY {
					plan = append(plan, [2]int{scannerIx, targetIx})
					known[targetIx] = true
					added = true
					continue
				}

				if existsY && !existsX {
					plan = append(plan, [2]int{targetIx, scannerIx})
					known[scannerIx] = true
					added = true
				}
			}
		}

		if !added {
			break
		}
	}

	return plan
}

type Scanner struct {
	pos         Coord
	orientation int
}

func collectBeacons(
	scannerReads *map[int][]*Coord,
	scanners *map[int]Scanner,
) *map[Coord]bool {
	beacons := make(map[Coord]bool)

	for ix, reads := range *scannerReads {
		scanner := (*scanners)[ix]
		ofn := Orientations[scanner.orientation]

		for _, read := range reads {
			coord := ofn(*read)
			beacons[Coord{
				x: coord.x + scanner.pos.x,
				y: coord.y + scanner.pos.y,
				z: coord.z + scanner.pos.z,
			}] = true
		}
	}

	return &beacons
}

func main() {
	scannerReads := parse(common.ReadLines(os.Args[1]))
	distances := make([]*[][]float64, len(scannerReads))
	for ix, reads := range scannerReads {
		distances[ix] = distanceMatrix(reads[:])
	}

	scanners := map[int]Scanner{
		0: Scanner{pos: Coord{x: 0, y: 0, z: 0}, orientation: 0},
	}

	for _, pair := range getPlan(&scannerReads, &distances) {
		scanners[pair[1]] = findScannerPosition(
			&scannerReads,
			&distances,
			pair[0], pair[1],
			scanners[pair[0]],
		)
	}

	fmt.Println("Part I", len(*collectBeacons(&scannerReads, &scanners)))

	maxDist := math.MinInt64

	for _, first := range scanners {
		for _, second := range scanners {
			dist := manhattan(&first.pos, &second.pos)

			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	fmt.Println("Part II", maxDist)
}
