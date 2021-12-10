package longestIncreasingSubsequence

import "fmt"

var cache map[string]int

func init() {
	cache = map[string]int{}
}

func solution(sequence []int) int {

	key := fmt.Sprint(sequence)
	if ans, ok := cache[key]; ok {
		return ans
	}

	l := len(sequence)

	// check if the sequence is strictly increasing
	var skipNone int
	end := sequence[0]
	for i := 1; i < l; i++ {
		if sequence[i] > end {
			skipNone++
			end = sequence[i]
		}
	}

	// fmt.Println("skip", sequence[1:])
	skip1st := solution(sequence[1:])

	// var next []int
	skip2nd := solution(append([]int{sequence[0]}, sequence[2:]...))

	var ans int
	for _, k := range []int{skipNone, skip1st, skip2nd} {
		if k > ans {
			ans = k
		}
	}

	cache[key] = ans
	return ans

}
