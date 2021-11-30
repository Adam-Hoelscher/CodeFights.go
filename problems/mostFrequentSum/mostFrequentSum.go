package mostFrequentSum

import "sort"

func mostFrequentSum(t *Tree) []int {

	counter := map[int]int{}

	var search func(t *Tree) int
	search = func(t *Tree) int {
		if t == nil {
			return 0
		}
		s := t.Value.(int)
		s += search(t.Left)
		s += search(t.Right)
		counter[s]++
		return s
	}

	search(t)

	sumsWithCount := map[int][]int{0: []int{}}
	var maxCount int

	for sum, count := range counter {
		sumsWithCount[count] = append(sumsWithCount[count], sum)
		if count > maxCount {
			maxCount = count
		}
	}

	sort.Ints(sumsWithCount[maxCount])

	return sumsWithCount[maxCount]

}
