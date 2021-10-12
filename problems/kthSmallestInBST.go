package problems

func kthSmallestInBST(t *Tree, k int) int {

	ch := make(chan interface{})
	for p := 0; p < k; p++ {
		<-ch
	}

	return (<-ch).(int)
}

func traverse(t *Tree, ch chan<- interface{}) {
	if t == nil {
		return
	}
	traverse(t.Left, ch)
	ch <- t.Value
	traverse(t.Right, ch)
}
