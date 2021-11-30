package largestValuesInTreeRows

import "math"

func largestValuesInTreeRows(t *Tree) []int {

	ans := []int{}

	var nextRow []*Tree

	for row := []*Tree{t}; len(row) > 0; row = nextRow {

		// establish the next queue
		nextRow = []*Tree{}

		// set the minimum to "negative infinity"
		max := math.MinInt64

		// for everything in the queu
		for _, t := range row {

			// if the node is null, move on
			if t == nil {
				continue
			}

			// update the maximum value
			val := t.Value.(int)
			if val > max {
				max = val
			}

			// add the children to the next queue
			nextRow = append(nextRow, t.Left)
			nextRow = append(nextRow, t.Right)
		}

		// if we had any values in this row, add to the answer
		if max > math.MinInt64 {
			ans = append(ans, max)
		}
	}

	return ans
}
