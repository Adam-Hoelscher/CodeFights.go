package problems

func deleteFromBST(t *Tree, queries []int) *Tree {

	for _, q := range queries {
		t = deleteSingle(t, q)
	}

	return t

}

func deleteSingle(t *Tree, value int) *Tree {
	if t == nil {
		return t
	} else if value < t.Value.(int) {
		t.Left = deleteSingle(t.Left, value)
	} else if value > t.Value.(int) {
		t.Right = deleteSingle(t.Right, value)
	} else if value == t.Value {
		if t.Left != nil {

			t.Value = findMax(t.Left)

			// The problem specifically says to "remove" the rightmost node of
			// the left child, which is apprently distinct from "deleting" that
			// value from the left child. Either one will result in a valid
			// BST, but the expected output of the tests assume the direct
			// removal, rather than a recursive delete.

			// t.Left = deleteSingle(t.Left, t.Value.(int))
			t.Left = removeRight(t.Left)

		} else {
			t = t.Right
		}
	}
	return t
}

func findMax(t *Tree) interface{} {
	for t.Right != nil {
		t = t.Right
	}
	return t.Value
}

func removeRight(t *Tree) *Tree {
	if t.Right == nil {
		return t.Left
	} else {
		t.Right = removeRight(t.Right)
		return t
	}
}
