package climbingStairs

var ClimbingStairs = solution

func solution(n int) int {

	cache := map[int]int{}

	var search func(int) int
	search = func(n int) int {

		key := n
		var v int

		if val, ok := cache[n]; ok {
			v = val

		} else {
			if n <= 2 {
				v = n

			} else {
				a := search(n - 2)
				b := search(n - 1)
				v = a + b
			}
		}

		cache[key] = v
		return v

	}

	return search(n)

}
