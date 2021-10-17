package problems

import "sort"

func findSubstrings(words []string, parts []string) []string {

	checks := map[int]map[byte][]string{}

	for _, p := range parts {
		l := len(p)
		key := byte(p[0])
		if _, ok := checks[l]; !ok {
			checks[l] = map[byte][]string{}
		}
		checks[l][key] = append(checks[l][key], p)
	}

	lensToCheck := []int{}
	for l := range checks {
		lensToCheck = append(lensToCheck, l)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lensToCheck)))

	answer := []string{}

	for _, w := range words {

	lengths:
		for _, l := range lensToCheck {
			for pos := 0; pos+l <= len(w); pos++ {
				tail := w[pos:]
				key := tail[0]
				for _, p := range checks[l][key] {
					newTail, change := searchWord(tail, p)
					if change {
						w = w[:pos] + newTail
						break lengths
					}
				}
			}
		}
		answer = append(answer, w)
	}

	return answer

}

func searchWord(word, part string) (string, bool) {
	l := len(part)
	if word[:l] == part {
		return "[" + part + "]" + word[l:], true
	} else {
		return word, false
	}
}
