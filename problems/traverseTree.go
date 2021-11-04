package problems

func traverseTree(t *Tree) []int {

	q := []*Tree{t}
	ans := []int{}

	for len(q) > 0 {

		t, q = q[0], q[1:]

		if t == nil {
			continue
		}

		ans = append(ans, t.Value.(int))

		q = append(q, t.Left)
		q = append(q, t.Right)

	}

	return ans

}
