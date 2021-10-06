package problems

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{}
	at := head

	for l1 != nil && l2 != nil {
		if l2.Value.(int) < l1.Value.(int) {
			at.Next = l2
			l2 = l2.Next
		} else {
			at.Next = l1
			l1 = l1.Next
		}
		at = at.Next
	}

	for l1 != nil {
		at.Next = l1
		l1 = l1.Next
		at = at.Next
	}

	for l2 != nil {
		at.Next = l2
		l2 = l2.Next
		at = at.Next
	}

	at.Next = nil

	return head.Next
}
