package composeRanges

import "fmt"

var ComposeRanges = solution

func solution(nums []int) []string {

	if len(nums) == 0 {
		return []string{}
	}

	runs := [][2]int{{nums[0], nums[0]}}

	for _, v := range nums {
		l := len(runs)
		if runs[l-1][1] >= v-1 {
			runs[l-1][1] = v
		} else {
			runs = append(runs, [2]int{v, v})
		}
	}

	ans := []string{}
	for _, pair := range runs {
		a := pair[0]
		b := pair[1]
		if a != b {
			ans = append(ans, fmt.Sprintf("%d->%d", a, b))
		} else {
			ans = append(ans, fmt.Sprint(a))
		}
	}
	return ans

}
