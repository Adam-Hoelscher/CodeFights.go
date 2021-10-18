package problems

type trie struct {
	children map[byte]*trie
	stop     bool
}

func newTrie() *trie {
	return &trie{
		children: map[byte]*trie{},
		stop:     false,
	}
}

func (t *trie) add(word string) {

	if len(word) > 0 {

		head := word[0]

		if _, ok := t.children[head]; !ok {
			t.children[head] = newTrie()
		}
		next := t.children[head]

		tail := word[1:]
		next.add(tail)

	} else {

		t.stop = true

	}
}

func (t *trie) search(word string) (int, bool) {

	if t == nil {
		return 0, false
	}

	if len(word) == 0 {
		return 0, t.stop
	}

	head := word[0]
	tail := word[1:]

	next := t.children[head]
	tailLen, found := next.search(tail)

	if found {
		return tailLen + 1, found
	} else {
		return 0, t.stop
	}

}

func findSubstrings(words []string, parts []string) []string {

	checks := newTrie()

	for _, p := range parts {
		checks.add(p)
	}

	for wIdx, w := range words {
		maxLen, maxIdx := 0, 0
		for i := 0; i < len(w); i++ {
			l, _ := checks.search(w[i:])
			if l > maxLen {
				maxLen, maxIdx = l, i
			}
		}
		if maxLen > 0 {
			before := w[:maxIdx]
			found := w[maxIdx : maxIdx+maxLen]
			after := w[maxIdx+maxLen:]
			words[wIdx] = before + "[" + found + "]" + after
		}
	}

	return words
}
