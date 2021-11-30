package rearrangeLastN

func rearrangeLastN(l *ListNode, n int) *ListNode {

	if l == nil {
		return l
	}

	oldHead := l
	newTail := l
	length := 1

	for l.Next != nil {
		l = l.Next
		length += 1
	}
	l.Next = oldHead

	l = oldHead
	for i := 0; i < 2*length-n-1; i++ {
		newTail = newTail.Next
	}

	answer := newTail.Next
	newTail.Next = nil

	return answer

}
