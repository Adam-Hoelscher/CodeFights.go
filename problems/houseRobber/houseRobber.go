package houseRobber

import "fmt"

var HouseRobber = solution

var cache map[string]int

func init() {
	cache = map[string]int{}
}

func solution(nums []int) int {

	key := fmt.Sprint(nums)
	var v int

	if val, ok := cache[key]; ok {
		v = val

	} else {
		if len(nums) == 0 {
			v = 0

		} else if len(nums) == 1 {
			v = nums[0]

		} else {
			take := nums[0] + solution(nums[2:])
			skip := solution(nums[1:])
			if take > skip {
				v = take
			} else {
				v = skip
			}
		}
	}

	cache[key] = v
	return v

}
