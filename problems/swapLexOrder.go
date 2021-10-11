package problems

func swapLexOrder(str string, pairs [][]int) string {

	length := len(str)

	// build a map from each position to the positions it could get a char from
	links := map[int][]int{}

	// each position can pull from itself
	for x := 0; x < length; x++ {
		// links[x] = []int{x}
	}

	// each position can also swap with anything it's linked to
	for _, p := range pairs {
		a, b := p[0]-1, p[1]-1
		links[a] = append(links[a], b)
		links[b] = append(links[b], a)
	}

	// each position is linked to some number of other positions
	// each group will get an ID equal to the lowest index of the group
	groupID := map[int]int{}
	// each group will have a pool of possible chars that they can access
	pools := map[int]map[byte]int{}

	// loop through all positions
	for i := 0; i < length; i++ {

		// check if we've already found a group for this position
		_, found := groupID[i]

		// if we didn't do DFS for the members of the group
		if !found {

			// while we are searching we can also build the pool
			pools[i] = map[byte]int{}

			// keep a set of positions we've visited to avoid cycles
			visited := map[int]bool{}

			// start the queue with the current position
			queue := []int{i}

			for len(queue) > 0 {

				at := queue[0]

				// if we haven't visited the position
				if !visited[at] {

					// add this character to the pool
					char := str[at]
					pools[i][char]++

					// add the all positions that are 1 away to the queue
					for _, next := range links[at] {
						queue = append(queue, next)
					}

					// mark that we've visited this position
					visited[at] = true

				}

				// remove the position we just checked from the queue
				queue = queue[1:]

				// mark each position we've visited as part of this group
				for j := range visited {
					groupID[j] = i
				}

			}
		}
	}

	// make an array of bytes for building the answer
	answer := make([]byte, length)

	// for each position in the
	for i := 0; i < length; i++ {

		// get the id of the position's group and the pool
		id := groupID[i]
		pool := pools[id]

		// find the maximum value left in the pool
		var v byte
		for char := range pools[id] {
			if char > v {
				v = char
			}
		}

		// put the max value in the current position
		answer[i] = v

		// decrement the count for the value we just used
		pools[id][v]--

		// if the count is now zero, remove the value from the pool
		if pool[v] == 0 {
			delete(pool, v)
		}

	}

	return string(answer)
}
