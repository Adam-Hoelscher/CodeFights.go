package firstDuplicate

func firstDuplicate(a []int) int {
	seen := map[int]bool{}
	for _, val := range a {
		exists := seen[val]
		if exists {
			return val
		} else {
			seen[val] = true
		}
	}
	return -1
}
