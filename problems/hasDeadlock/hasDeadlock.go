package hasDeadlock

import "container/list"

// basically, we have a directed graph and we're looking for cycles
func hasDeadlock(connections [][]int) bool {

	// var search func(int, map[int]bool, map[int]bool) bool
	// search = func(src int, used, done map[int]bool) bool {
	// }

	for src := range connections {

		q := list.New()
		q.PushBack(src)

		seen := map[int]bool{}

		for q.Front() != nil {

			at := q.Remove(q.Front()).(int)

			if at == src && seen[at] {
				return true
			}

			if seen[at] {
				continue
			}

			seen[at] = true
			for _, dst := range connections[at] {
				q.PushBack(dst)
			}
		}

	}

	return false
}
