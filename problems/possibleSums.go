package problems

func possibleSums(coins []int, quantity []int) int {

	sums := map[int]bool{0: true}

	for idx, val := range coins {
		number := quantity[idx]
		newSums := map[int]bool{0: true}
		for base := range sums {
			for use := 0; use <= number; use++ {
				newVal := base + use*val
				newSums[newVal] = true
			}
		}
		sums = newSums
	}

	return len(sums) - 1

}
