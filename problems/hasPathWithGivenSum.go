package problems

func hasPathWithGivenSum(t *Tree, s int) bool {

	if t == nil {
		return false
	}

	s -= t.Value.(int)

	if t.Left == nil && t.Right == nil {
		return s == 0
	} else {
		return hasPathWithGivenSum(t.Left, s) || hasPathWithGivenSum(t.Right, s)
	}

}
