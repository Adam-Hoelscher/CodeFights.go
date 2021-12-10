import "unicode"

func solution(s string) string {

	ans := ""

	for _, char := range s {
		if unicode.IsUpper(char) {
			if ans != "" {
				ans += " "
			}
			ans += string(unicode.ToLower(char))
		} else {
			ans += string(char)
		}
	}

	return ans

}
