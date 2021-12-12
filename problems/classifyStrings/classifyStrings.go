package classifystrings

var ClassifyStrings = solution

var vowels map[rune]bool = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

var cache map[string]string = map[string]string{}

func solution(s string) string {

	if ans, ok := cache[s]; ok {
		return ans
	}

	for idx, char := range s {
		if char == '?' {
			a := solution(s[:idx] + "a" + s[idx+1:])
			b := solution(s[:idx] + "b" + s[idx+1:])
			if a == b {
				return a
			} else {
				return "mixed"
			}
		}
	}

	var vCount, cCount int

	ans := "good"
	for _, char := range s {
		if vowels[char] {
			vCount++
			cCount = 0
		} else {
			cCount++
			vCount = 0
		}
		if vCount > 2 || cCount > 4 {
			ans = "bad"
			break
		}
	}

	cache[s] = ans
	return ans

}
