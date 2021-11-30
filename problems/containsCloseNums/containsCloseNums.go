package containsCloseNums

func containsCloseNums(nums []int, k int) bool {

	last := map[int]int{}

	for idx, val := range nums {
		lastIdx, ok := last[val]
		if ok && idx-lastIdx <= k {
			return true
		} else {
			last[val] = idx
		}
	}

	return false
}
