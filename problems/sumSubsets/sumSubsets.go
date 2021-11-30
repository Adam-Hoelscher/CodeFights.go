package sumSubsets

import "fmt"

func sumSubsets(arr []int, num int) [][]int {

	// if the target is zero we can definitely get there with an empty list
	if num == 0 {
		return [][]int{{}}
	} else

	// if the array is empty or the target is negative there is no answer
	if len(arr) == 0 || num < 0 {
		return [][]int{}
	}

	// make a list of possible answers and a set of answer's we seen
	// the set will use Sprint strings as keys
	ans := [][]int{}
	seen := map[string]bool{}

	// for each item in the array
	for i, n := range arr {

		// make a list of ways we could make a smaller target after removing
		// this value from the sum
		tails := sumSubsets(arr[i+1:], num-n)

		// for each of those tails
		for _, t := range tails {

			// put together the full subset and the key for the set
			subset := append([]int{n}, t...)
			key := fmt.Sprint(subset)

			// if we've not seen this answer before, add it
			if !seen[key] {
				ans = append(ans, subset)
				seen[key] = true
			}
		}
	}

	return ans
}
