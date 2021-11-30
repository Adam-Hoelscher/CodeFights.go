package graphDistances

import (
	"container/heap"
)

type graphPath struct {
	location int
	length   int
}

type pathHeap []*graphPath

func (h pathHeap) Len() int           { return len(h) }
func (h pathHeap) Less(i, j int) bool { return h[i].length < h[j].length }
func (h pathHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pathHeap) Push(x interface{}) {
	*h = append(*h, x.(*graphPath))
}

func (h *pathHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func graphDistances(g [][]int, s int) []int {

	// make an array to hold the answer
	minDist := make([]int, len(g))

	// make a set to hold vertices we've expanded
	seen := map[int]bool{}

	// initializae the queue with the null path
	q := pathHeap{&graphPath{s, 0}}

	// while there are paths in the heap
	for len(q) > 0 {

		// get the shortest path from the heap
		at := heap.Pop(&q).(*graphPath)

		// if we've already expanded this vertex skip it
		if seen[at.location] {
			continue
		}

		// mark that we've expanded the vertex
		seen[at.location] = true

		// note the distance
		minDist[at.location] = at.length

		// for each vertex in the graph
		for dst, cost := range g[at.location] {

			// if it's not connected move on
			if cost == -1 {
				continue
			}

			// the next path adds the weight of this edge
			next := graphPath{
				location: dst,
				length:   at.length + cost,
			}

			// add the new path to the heap
			heap.Push(&q, &next)
		}

	}

	return minDist

}
