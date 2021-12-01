package maximalSquare

var MaximalSquare = solution

func solution(matrix [][]string) int {

	rCount := len(matrix)
	if rCount == 0 {
		return 0
	}

	cCount := len(matrix[0])

	max := 0

	for i, row := range matrix {
		for j, val := range row {

			for size, valid := 1, val == "1"; valid; size++ {

				maxStep := size - 1

				if i+maxStep == rCount || j+maxStep == cCount {
					valid = false
					continue
				}

				for step := 0; step <= maxStep && valid; step++ {

					if matrix[i+maxStep][j+step] == "0" {
						valid = false
					}

					if matrix[i+step][j+maxStep] == "0" {
						valid = false
					}

				}

				if matrix[i+maxStep][j+maxStep] == "0" {
					valid = false
				}

				if valid && size > max {
					max = size
				}
			}

		}
	}

	return max * max

}
