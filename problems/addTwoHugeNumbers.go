package problems

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {

	reverseList := func(l *ListNode) *ListNode {
		var tail *ListNode = nil
		for l != nil {
			trueNext := l.Next
			l.Next = tail
			tail = l
			l = trueNext
		}
		return tail
	}

	a = reverseList(a)
	b = reverseList(b)

	var answer *ListNode = nil
	carry := 0

	for a != nil || b != nil || carry != 0 {

		nodeValue := carry

		if a != nil {
			nodeValue += a.Value.(int)
			a = a.Next
		}

		if b != nil {
			nodeValue += b.Value.(int)
			b = b.Next
		}

		carry = nodeValue / 10_000
		nodeValue %= 10_000

		answer = &ListNode{
			Value: nodeValue,
			Next:  answer,
		}
	}

	return answer
}
