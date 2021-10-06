package problems

func reverseNodesInKGroups(l *ListNode, k int) *ListNode {

	// count elements so we know how many groups we can reverse
	remaining := 0
	for c := l; c != nil; c = c.Next {
		remaining++
	}

	head := &ListNode{}
	priorTail := head

	for remaining > 0 {

		var groupHead *ListNode = nil
		var groupTail *ListNode = l

		if remaining >= k {

			for i := 0; i < k; i++ {
				trueNext := l.Next
				l.Next = groupHead
				groupHead = l
				l = trueNext
			}

		} else {
			// not big enough to reverse
			groupHead = l
		}

		priorTail.Next = groupHead
		priorTail = groupTail
		remaining -= k
	}

	return head.Next
}
