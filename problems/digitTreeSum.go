package problems

func digitTreeSum(t *Tree) int64 {

	var rf func(t *Tree, i int) int
	rf = func(t *Tree, i int) int {

		var sum int
		i += t.Value.(int)

		if t.Left != nil || t.Right != nil {

			if t.Left != nil {
				sum += rf(t.Left, 10*i)
			}
			if t.Right != nil {
				sum += rf(t.Right, 10*i)
			}

		} else {

			sum = i

		}

		return sum

	}

	return int64(rf(t, 0))

}
