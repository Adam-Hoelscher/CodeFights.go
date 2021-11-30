package nearestGreater

func nearestGreater(a []int) []int {

	var stack []priorNode

	// make a slice for the answer since we know the length
	ans := make([]int, len(a))

	// create stackd for nodes that need updated
	stack = []priorNode{}

	// build the ans list
	for i, v := range a {

		// start by assuming that there are no larger values
		ans[i] = -1

		// stack invariant: values on the stack are non-increasing;
		// a[i+1] cannot go on the stack before a[i]
		// if a[i+1] > a[i] then a[i+1] will clear a[i] from the stack
		for len(stack) > 0 && v > stack[len(stack)-1].value {
			// pop from the stack and save the current value as the answer for
			// the prior node
			prior := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prior.index] = i
		}

		// add a priorNode for a[i]
		newPrior := priorNode{i, v}
		stack = append(stack, newPrior)
	}

	// reset the stack and iterate the other way
	stack = []priorNode{}

	for i := len(a) - 1; i >= 0; i-- {
		v := a[i]

		for len(stack) > 0 && v > stack[len(stack)-1].value {

			prior := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// figure out if the highger value to the left is better
			lDist := prior.index - i
			rDist := ans[prior.index] - prior.index

			if ans[prior.index] == -1 || lDist <= rDist {
				ans[prior.index] = i
			}
		}

		newPrior := priorNode{i, v}
		stack = append(stack, newPrior)
	}

	return ans
}
