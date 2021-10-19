package problems

// basically, we have a directed graph and we're looking for cycles
func hasDeadlock(connections [][]int) bool {

	// for each place we can start
	for src, q := range connections {

		// keep track of where we've bee
		seen := map[int]bool{}

		// while we have a queue
		for len(q) > 0 {

			// pull from front of queue
			at := q[0]
			q = q[1:]

			// if we're back where we started, we found a cycle
			if at == src {
				return true

				// if we've not expanded this node, do so
			} else if !seen[at] {
				seen[at] = true
				q = append(q, connections[at]...)
			}
		}
	}

	// if we didn't find a cycle then there aren't any
	return false
}
