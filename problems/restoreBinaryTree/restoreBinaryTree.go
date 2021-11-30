package restoreBinaryTree

func restoreBinaryTree(inorder []int, preorder []int) *Tree {

	if len(inorder) == 0 {
		return nil
	}

	// the root of the tree must be the first value of preorder
	rootVal := preorder[0]

	// find the position of the root value in the inorder
	var pos int
	for pos = 0; inorder[pos] != rootVal; pos++ {
	}

	// remove the root both lists
	newPre := preorder[1:]
	newIn := append(inorder[:pos], inorder[pos+1:]...)

	// everything before the root is the left child
	lIn := newIn[:pos]
	lPre := newPre[:pos]
	lChild := restoreBinaryTree(lIn, lPre)

	// everything after the root is the right child
	rIn := newIn[pos:]
	rPre := newPre[pos:]
	rChild := restoreBinaryTree(rIn, rPre)

	return &Tree{rootVal, lChild, rChild}

}
