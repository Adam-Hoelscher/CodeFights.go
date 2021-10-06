package problems

func rotateImage(a [][]int) [][]int {

	var l = len(a[0])

	// IDOIMATIC: space O(n^2), time O(n^2)
	// output := make([][]int, l)
	// for row := range output {
	//     output[row] = make([]int, l)
	// }

	// for i := range a {
	//     for j := range a {
	//         output[j][l-i-1] = a[i][j]
	//     }
	// }

	// return output

	// SPACE EFFICIENT: space O(1), time O(n^2)
	// rotating in place can be done with a combination of transposing and flipping
	// right rotation = low-left/up-right transpose + horizontal flip
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			// if j is greater than i, we are above the diagonal and need to transpose
			if j >= i {
				swap(a, i, j, j, i)
			}
			// if j is more than 1/2 of i, we are past the middle and need to flip
			if j >= l/2 {
				swap(a, i, j, i, l-j-1)
			}
		}
	}

	return a
}

func swap(a [][]int, i, j, m, n int) {
	a[i][j], a[m][n] = a[m][n], a[i][j]
}
