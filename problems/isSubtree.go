package problems

func isSubtree(t1 *Tree, t2 *Tree) bool {
	if t2 == nil {
		return true
	} else if t1 == nil {
		return false
	} else {
		ret := isEqualTree(t1, t2) ||
			isSubtree(t1.Left, t2) ||
			isSubtree(t1.Right, t2)
		return ret
	}
}

func isEqualTree(t1 *Tree, t2 *Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil || t1.Value != t2.Value {
		return false
	} else {
		return isEqualTree(t1.Left, t2.Left) && isEqualTree(t1.Right, t2.Right)
	}
}
