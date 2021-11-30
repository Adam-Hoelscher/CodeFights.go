package networkWires

import "sort"

func networkWires(n int, wires [][]int) int {

	// add variable to hold the answe
	ans := 0

	// sort the wires so we do the shortest first
	sort.Slice(wires, func(i, j int) bool {
		return wires[i][2] < wires[j][2]
	})

	// make a map of the components
	components := map[int]*map[int]bool{}
	// start with each component being one node
	for i := 0; i < n; i++ {
		components[i] = &map[int]bool{i: true}
	}

	// for each wire
	for _, w := range wires {

		// see what component it's part of
		aComp := components[w[0]]
		bComp := components[w[1]]

		// if they're already in the same component, move on
		if aComp == bComp {
			continue
		}

		// add this wire's length to join the components
		ans += w[2]

		// merge component a and component b
		// for each node in component b
		for k := range *bComp {
			// add the node to component a
			(*aComp)[k] = true
			// tag node with component a
			components[k] = aComp
		}

	}

	return ans

}
