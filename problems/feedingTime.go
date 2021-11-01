package problems

func feedingTime(classes [][]string) int {

	// set the limit for the number of colors
	colorLimit := 6

	// find the size of the graph
	l := len(classes)

	// keep track of which classes want to feed which animals
	feedings := map[string][]int{}

	// keep track of which nodes are connected
	graph := make([]map[int]bool, l)
	// start with each class having no connections
	for idx := range classes {
		graph[idx] = map[int]bool{}
	}

	// for each class
	for idx, c := range classes {
		// for each animal the class wants to feed
		for _, animal := range c {
			// for each other class that wants to feed the anmial
			for _, n := range feedings[animal] {
				// connect the 2 classes
				graph[idx][n] = true
				graph[n][idx] = true
			}
			// record that this class wants to feed the animal
			feedings[animal] = append(feedings[animal], idx)
		}
	}

	// hold the answer in a variable
	var ans int

	// start with 1 color and search until we hit the limt
	for ans = 1; ans < colorLimit; ans++ {

		// make up possible colors
		colors := make([]int, l)

		// color the first node with the first color
		colors[0] = 1

		// try coloring each node starting with the 2nd
		var idx int
		for idx = 1; 0 < idx && idx < l; {

			// if we've not on the last option, try the next
			if colors[idx] < ans {

				// move this node to the next color
				colors[idx]++

				// move to the next if we don't have a problem
				if checkColors(graph, colors) {
					idx++
				}

			} else {
				// reset the current node to no color
				colors[idx] = 0
				// back up to the last node
				idx--
			}

		}

		// if we made it to the end, we can stop searching
		if idx == l {
			break
		}
	}

	// if we hit the limit, return a -1
	if ans == colorLimit {
		ans = -1
	}

	return ans

}

func checkColors(graph []map[int]bool, colors []int) bool {

	// for each node in the graph
	for src, connections := range graph {

		// if we haven't colored the node, it doesn't matter yet
		if colors[src] == 0 {
			continue
		}

		for dst := range connections {
			// if 2 connected nodes have the same color, this isn't valid
			if colors[src] == colors[dst] {
				return false
			}
		}
	}

	// if we didn't find a problem, this *is* valid
	return true
}
