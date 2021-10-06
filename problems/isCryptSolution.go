package problems

import "strconv"

func isCryptSolution(crypt []string, solution [][]string) bool {

	lookup := map[string]string{}
	for _, pair := range solution {
		key := pair[0]
		val := pair[1]
		lookup[key] = val
	}

	valueStrings := [3]string{}
	for wordNum, word := range crypt {
		for _, char := range word {
			valueStrings[wordNum] += lookup[string(char)]
		}
	}

	values := [3]int{}
	for i, str := range valueStrings {
		values[i], _ = strconv.Atoi(str)
		// check for leading zeros
		if strconv.Itoa(values[i]) != str {
			return false
		}
	}

	return values[0]+values[1] == values[2]
}
