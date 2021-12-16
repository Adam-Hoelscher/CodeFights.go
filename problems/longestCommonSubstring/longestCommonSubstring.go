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

func (t *trie) search(word string) int {

	if t == nil {
		return 0
	}

	if len(word) == 0 {
		return 0
	}

	head, tail := word[0], word[1:]
	next := t.children[head]    

	if next != nil {
		return 1 + next.search(tail)
	} else {
		return 0
	}
}

func solution(s string, t string) int {

    trie := newTrie()
    for i := 0; i < len(s); i++ {
        trie.add(s[i:])
    }

    ans := 0
    for i := 0; i + ans < len(t); i++ {
        l := trie.search(t[i:])
        if l > ans {
            ans = l
        }
    }

    return ans

}
