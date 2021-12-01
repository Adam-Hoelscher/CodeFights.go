package regularExpressionMapping

var regularExpressionMapping = solution

func solution(s string, p string) bool {

	// if the pattern is exhausted only the empty string is valid
	if len(p) == 0 {
		return len(s) == 0
	}

	// decide if the pattern starts with a starred token
	starToken := len(p) > 1 && p[1] == '*'

	// decide if the first character matches the first token
	headMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	// if the pattern starts with a starred token
	if starToken {

		// we can  try matching without the token
		if solution(s, p[2:]) {
			return true
		}

	} else {
		// we need to consume a token
		p = p[1:]

	}

	// if the head matches and the rest is valid
	return headMatch && solution(s[1:], p)

}
