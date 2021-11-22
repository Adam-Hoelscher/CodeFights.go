package problems

import (
	"fmt"
	"sort"
)

func solution(a []int, sum int) string {

	var find func([]int, int) [][]int
	find = func(a []int, sum int) [][]int {

		if sum < 0 {
			return [][]int{}
		}

		if sum == 0 {
			return [][]int{{}}
		}

		ans := [][]int{}
		for idx, val := range a {
			for _, tail := range find(a[idx:], sum-val) {
				head := []int{val}
				ans = append(ans, append(head, tail...))
			}
		}
		return ans
	}

	// deduplicate and sort the arr
	set := map[int]bool{}
	dedupe := []int{}
	for _, v := range a {
		if !set[v] {
			dedupe = append(dedupe, v)
		}
		set[v] = true
	}
	sort.Ints(dedupe)

	ansInts := find(dedupe, sum)

	ansStr := ""

	for _, arr := range ansInts {
		ansStr += "("
		for idx, val := range arr {
			if idx != 0 {
				ansStr += " "
			}
			ansStr += fmt.Sprint(val)
		}
		ansStr += ")"
	}

	if ansStr == "" {
		ansStr = "Empty"
	}

	return ansStr
}
