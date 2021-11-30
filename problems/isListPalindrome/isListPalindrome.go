package isListPalindrome

func isListPalindrome(l *ListNode) bool {

	head := l

	// get the length of the list
	length := 0
	for l != nil {
		length++
		l = l.Next
	}

	// move to the middle of the list
	at := head
	for i := 0; i < length/2; i++ {
		at = at.Next
	}

	var tail *ListNode = nil
	// move through the rest of the list reversing the links
	for at != nil {
		trueNext := at.Next
		at.Next = tail
		tail = at
		at = trueNext
	}

	// peel off pairs of head and tail and stop if no match
	for head != nil && tail != nil {
		if head.Value != tail.Value {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}
