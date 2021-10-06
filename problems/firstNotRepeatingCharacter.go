package problems

func firstNotRepeatingCharacter(s string) string {

	counter := map[rune]int{}

	for _, char := range s {
		counter[char] += 1
	}

	for _, char := range s {
		if counter[char] < 2 {
			return string(char)
		}
	}

	return "_"
}
