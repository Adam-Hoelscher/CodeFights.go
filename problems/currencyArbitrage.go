package problems

func currencyArbitrage(exchange [][]float64) bool {

	l := len(exchange)

	// Use the same logic as Bellman-Ford
	highest := make([]float64, l)
	highest[0] = 1

	// Bellman-Ford needs a maximum of N-1 relaxations but we'll stop early
	// if we stop relaxing
	for i, relaxing := 0, true; i < l && relaxing; i, relaxing = i+1, false {

		// for each source
		for u := 0; u < l; u++ {
			// for each destination
			for v := 0; v < l; v++ {

				// calculate the amount if we go to u and then v
				rate := exchange[u][v]
				alt := highest[u] * rate

				// if that's better than the current amount for v replace it
				if alt > highest[v] {
					highest[v] = alt
					relaxing = true
				}
			}
		}
	}

	return highest[0] > 1

}
