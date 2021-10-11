package problems

// Singly-linked lists are already defined with this interface:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

// Binary trees are already defined with this interface:
type Tree struct {
	Value interface{}
	Left  *Tree
	Right *Tree
}
