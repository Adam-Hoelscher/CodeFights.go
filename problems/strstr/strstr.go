package strstr

var Strstr = solution

func solution(haystack string, needle string) int {

	// if the needle is the empty string, it necessarily matches at position 0
	if needle == "" {
		return 0
	}

	// get the lengths of both strings; we'll need them
	nLen := len(needle)
	hLen := len(haystack)

	// PREPROCESSING
	// make an memo of the longest prefix/suffix of each position in the needle
	memo := make([]int, nLen)
	for nIdx, mIdx := 1, 0; nIdx < nLen; {

		if needle[nIdx] == needle[mIdx] {
			// if the new character matches the current prefix extend the lps
			mIdx++
			memo[nIdx] = mIdx

			// move on to the next character in the needle
			nIdx++

		} else if mIdx == 0 {
			// if there is no current prefix move to the next character
			memo[nIdx] = 0
			nIdx++

		} else {
			// otherwise backtrack to the next shortest prefix
			mIdx = memo[mIdx-1]
		}
	}

	// SEARCH
	// check each position in the haystack for the needle
	for hIdx, nIdx := 0, 0; hIdx < hLen; {

		if haystack[hIdx] == needle[nIdx] {

			// if we have a match move forward
			nIdx++
			hIdx++

			// if we're past the length of the needle, we found a match
			if nIdx == nLen {
				return hIdx - nIdx
			}

		} else {

			if nIdx == 0 {
				// if we've checked the whole memo move to the next character
				hIdx++

			} else {
				// otherwise step back in the memo
				nIdx = memo[nIdx-1]
			}
		}
	}

	// if we passed the end of the haystack, the needle isn't present
	return -1

}
