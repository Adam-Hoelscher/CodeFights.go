package problems

func isTreeSymmetric(t *Tree) bool {

	// set up queues for comparison
	lSide := []*Tree{t}
	rSide := []*Tree{t}

	// while either of the queues has items
	for len(lSide) > 0 && len(rSide) > 0 {

		// pull from the queues
		lTree := lSide[0]
		lSide = lSide[1:]
		rTree := rSide[0]
		rSide = rSide[1:]

		if lTree == nil && rTree == nil {
			// if both are nil continue
			continue
		} else if lTree == nil || rTree == nil {
			// if only one is nil the branches don't match
			return false
		} else if lTree.Value != rTree.Value {
			// if the values aren't the same the branches don't match
			return false
		}

		// add the outer child and then the inner child to each queue
		lSide = append(lSide, lTree.Left, lTree.Right)
		rSide = append(rSide, rTree.Right, rTree.Left)

	}

	// if either queue is not empty then the branches don't match
	return len(lSide)+len(rSide) == 0

}
