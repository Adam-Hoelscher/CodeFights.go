package problems

import "sort"

type bogglePoint struct {
	X int
	Y int
}

func solution(board [][]string, words []string) []string {

	// make a map of each letter to the locations where it can be found
	locations := map[string][]bogglePoint{}
	for rowNum, row := range board {
		for colNum, letter := range row {
			p := bogglePoint{rowNum, colNum}
			locations[letter] = append(locations[letter], p)
		}
	}

	// define helper function that returns a map of the 8 neighbors of a point
	getNeighbors := func(p bogglePoint) map[bogglePoint]bool {
		ans := map[bogglePoint]bool{}
		steps := [3]int{-1, 0, 1}
		for _, dx := range steps {
			for _, dy := range steps {
				if !(dx == 0 && dy == 0) {
					ans[bogglePoint{p.X + dx, p.Y + dy}] = true
				}
			}
		}
		return ans
	}

	// define recursive function to search for a single word
	var checkOne func(string, map[bogglePoint]bool, *bogglePoint) bool
	checkOne = func(
		word string,
		used map[bogglePoint]bool,
		last *bogglePoint) bool {

		// if we're searching for an empty string, we're guaranteed to find it
		if len(word) == 0 {
			return true
		}

		// break the word into the first letter and the rest
		first, rest := word[:1], word[1:]

		// find the neighbors of the last point
		var neighbors map[bogglePoint]bool
		// if last is nil then we just started and can use any letter
		if last != nil {
			neighbors = getNeighbors(*last)
		}

		// for each location where this letter appears
		for _, loc := range locations[first] {

			validStep := (last == nil || neighbors[loc])

			// if we haven't been to this location and it's a valid step
			if !used[loc] && validStep {

				// add the new location to the used set
				used[loc] = true

				// if we can make the rest of the word using this first
				if checkOne(rest, used, &loc) {
					return true
				}

				// remove the new location for the next iteration
				delete(used, loc)
			}
		}

		// we didn't find the word
		return false
	}

	ans := []string{}
	// for each word check if we can find it and add it to the answer if we can
	for _, w := range words {
		if checkOne(w, map[bogglePoint]bool{}, nil) {
			ans = append(ans, w)
		}
	}

	// sort the answer slice
	sort.Strings(ans)

	return ans

}
