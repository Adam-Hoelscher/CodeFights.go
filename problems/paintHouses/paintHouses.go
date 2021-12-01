package paintHouses

var PaintHouses = solution

func solution(cost [][]int) int {

	l := len(cost)
	minArray := make([][3]int, l)

	// for the first house, copy the marginal as the total
	copy(minArray[0][:], cost[0])

	// for each house after the first
	for h := 1; h < l; h++ {
		// for each color we could paint the house
		for c := 0; c < 3; c++ {

			// look at the totals from the prior house
			priorHouse := minArray[h-1]

			// get the min total from each valid color for the prior house
			prior1 := priorHouse[(c+1)%3]
			prior2 := priorHouse[(c+2)%3]

			// record total cost from whichever valid color is cheaper
			if prior1 < prior2 {
				minArray[h][c] = prior1
			} else {
				minArray[h][c] = prior2
			}

			// add in the cost to paint this house this color
			minArray[h][c] += cost[h][c]
		}
	}

	// find the minimum total across all colors at the last house
	lastHouse := minArray[l-1]
	min := lastHouse[0]

	for c := 1; c < 3; c++ {
		if lastHouse[c] < min {
			min = lastHouse[c]
		}
	}

	return min

}
