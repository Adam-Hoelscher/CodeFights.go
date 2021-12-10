package kpalindrome

func solution(s string, k int) bool {

	l := len(s)

	if l < 2 {
		// if the string is <2 characters it's necessarily a palindrome
		return true
	}

	if s[0] == s[l-1] {
		// if the front and the back match, there's nothing gained by removing
		// either of them
		return solution(s[1:l-1], k)
	}

	if k == 0 {
		// if it's not already a palindrome and we have no removals left, we're
		// stuck
		return false
	}

	// otherwise we might be able to find an answer by spending one removal at
	// the front or the back
	return solution(s[1:], k-1) || solution(s[:l-1], k-1)

}
