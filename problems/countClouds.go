package problems

import "fmt"

type point struct {
	x int
	y int
}

func countClouds(skyMap [][]string) int {

	clouds := 0
	seen := map[point]bool{}
	for x := 0; x < len(skyMap); x++ {
		for y := 0; y < len(skyMap[x]); y++ {
			p := point{x, y}
			if search(p, skyMap, seen) {
				clouds++
				fmt.Println()
			}
		}
	}
	return clouds
}

func search(p point, skyMap [][]string, seen map[point]bool) bool {
	h := len(skyMap)
	w := len(skyMap[0])
	x := p.x
	y := p.y
	if x < 0 || y < 0 || x >= h || y >= w {
		return false
	} else if seen[p] {
		return false
	} else if skyMap[x][y] == "0" {
		return false
	} else {
		seen[p] = true
		for dx := -1; dx <= 1; dx++ {
			nextP := point{x + dx, y}
			search(nextP, skyMap, seen)
		}
		for dy := -1; dy <= 1; dy++ {
			nextP := point{x, y + dy}
			search(nextP, skyMap, seen)
		}
		return skyMap[x][y] == "1"
	}
}
