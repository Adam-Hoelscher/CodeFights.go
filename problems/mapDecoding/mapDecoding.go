package mapDecoding

import "strconv"

var MapDecoding = solution

const mod int = int(1e9 + 7)

var cache map[string]int

func init() {
	cache = map[string]int{}
}

func solution(message string) int {

	if ans, ok := cache[message]; ok {
		return ans
	}

	if len(message) == 0 {
		return 1
	}

	if message[0] == '0' {
		return 0
	}

	var take1, take2 int

	take1 = solution(message[1:])

	if len(message) > 1 {
		value, _ := strconv.Atoi(string(message[:2]))
		if value < 27 {
			take2 = solution(message[2:])
		} else {
			take2 = 0
		}
	}

	ans := (take1 + take2) % mod

	cache[message] = ans
	return ans

}
