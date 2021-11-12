package problems

func climbingStaircase(n int, k int) [][]int {

	// if we've no steps return a single empty path
	if n == 0 {
		return [][]int{{}}
	}

	// make an empty list for the answer
	ans := [][]int{}

	// for each jump size we could make
	for s := 1; s <= k && s <= n; s++ {

		// get all the ways we can finish
		tails := climbingStaircase(n-s, k)

		// for each way we could finish
		for _, t := range tails {

			// build the full path
			path := append([]int{s}, t...)

			// add this path to the list
			ans = append(ans, path)
		}
	}

	return ans

}
