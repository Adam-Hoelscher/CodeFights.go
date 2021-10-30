package problems

func nextLarger(a []int) []int {

	// make a slice for the answer since we know the length
	ans := make([]int, len(a))

	// create a stack for nodes that need updated
	stack := []priorNode{}

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
			ans[prior.index] = v
		}

		// add a priorNode for a[i]
		newPrior := priorNode{i, v}
		stack = append(stack, newPrior)
	}

	return ans
}
