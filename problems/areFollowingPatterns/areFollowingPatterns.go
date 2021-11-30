package areFollowingPatterns

func areFollowingPatterns(strings []string, patterns []string) bool {

	sLast := map[string]int{}
	pLast := map[string]int{}

	for idx, s := range strings {
		p := patterns[idx]
		sIdx, sOk := sLast[s]
		pIdx, pOk := pLast[p]
		if (sOk != pOk) || (sIdx != pIdx) {
			return false
		} else {
			sLast[s] = idx
			pLast[p] = idx
		}
	}

	return true

}
