package kpalindrome

func solution(s string, k int) bool {

	l := len(s)

	if l < 2 {
		return true
	}

	if s[0] == s[l-1] {
		return solution(s[1:l-1], k)
	}

	if k == 0 {
		return false
	}

	return solution(s[1:], k-1) || solution(s[:l-1], k-1)

}
