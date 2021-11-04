package problems

type spofPath struct {
	steps []int
	times map[int]int
}

func (p spofPath) extend(loc int) spofPath {

	newTimes := map[int]int{}
	for k, v := range p.times {
		newTimes[k] = v
	}

	if l := len(p.steps); l > 0 {
		newTimes[p.getPosition()] = l - 1
	}

	l := len(p.steps)
	newSteps := make([]int, l+1)
	copy(newSteps[:l], p.steps)
	newSteps[l] = loc

	return spofPath{steps: newSteps, times: newTimes}

}

func (p spofPath) findTime(loc int) int {
	return p.times[loc]
}

func (p spofPath) getTail(time int) []int {
	l := len(p.steps)
	return p.steps[time:l]
}

func (p spofPath) getPosition() int {
	l := len(p.steps)
	return p.steps[l-1]
}

func (p spofPath) getPrior() *int {
	l := len(p.steps)
	if l > 1 {
		return &p.steps[l-2]
	} else {
		return nil
	}
}

func (p spofPath) hasCycle() bool {
	l := len(p.steps)
	tail := p.steps[l-1]
	_, ok := p.times[tail]
	return ok
}

func newPath() spofPath {
	return spofPath{steps: []int{}, times: map[int]int{}}
}

func SinglePointOfFailure(connections [][]int) int {

	// assume everything is a spof
	spof := map[int]bool{}
	for i := range connections {
		spof[i] = true
	}

	// build a queue starting at position 0
	firstNode := newPath().extend(0)
	q := []spofPath{firstNode}

	// bfs
	for len(q) > 0 {

		at := q[0]
		q = q[1:]

		position := at.getPosition()

		// if we just made a cycle or if we already know this isn't spof
		// mark everything in the cycle as not spof
		if at.hasCycle() || !spof[position] {
			time := at.findTime(position)
			for _, p := range at.getTail(time) {
				if p != position {
					delete(spof, p)
				}
			}
		} else

		// otherwise keep going
		{
			for dst, edge := range connections[position] {

				connected := edge == 1

				prior := at.getPrior()
				uTurn := prior != nil && *prior == dst

				if connected && !uTurn {
					next := at.extend(dst)
					q = append(q, next)
				}
			}
		}
	}

	// make sure 0 isn't marked as spof, because it doesn't represent an edge
	delete(spof, 0)
	return len(spof)
}
