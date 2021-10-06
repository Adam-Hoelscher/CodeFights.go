package problems

func removeKFromList(l *ListNode, k int) *ListNode {

	head := &ListNode{Next: l}

	for at := head; at != nil; at = at.Next {
		temp := at.Next
		for temp != nil && temp.Value == k {
			temp = temp.Next
		}
		at.Next = temp
	}

	return head.Next
}
