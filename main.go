// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

type queueItem struct {
	priority int
	value    interface{}
}

type queueHeap []*queueItem

func (pq *queueHeap) Push(x interface{}) {
	item := x.(*queueItem)
	*pq = append(*pq, item)
}

func (pq *queueHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item.value
}

func (pq queueHeap) Len() int { return len(pq) }

func (pq queueHeap) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq queueHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

type PriorityQueue struct {
	heap queueHeap
}

func (pq PriorityQueue) Empty() bool { return pq.Len() == 0 }

func (pq PriorityQueue) Len() int { return len(pq.heap) }

func (pq *PriorityQueue) Push(priortiy int, value interface{}) queueItem {
	item := queueItem{priority: priortiy, value: value}
	heap.Push(&pq.heap, &item)
	return item
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(&pq.heap)
}

func main() {

	pq := PriorityQueue{}

	stuff := [][2]interface{}{
		{7, "a"},
		{0, "e"},
		{1, "g"},
		{2, "b"},
		{3, "c"},
		{6, "d"},
		{4, "f"},
		{5, "h"},
	}

	for _, i := range stuff {

		p := i[0].(int)
		v := i[1]

		fmt.Println("Add", pq.Push(p, v))
	}

	fmt.Println()
	for !pq.Empty() {
		fmt.Println("Pop", pq.Pop())
	}
}
