package nQueens

func nQueens(n int) [][]int {

	// keep a list of answers
	ans := [][]int{}

	// helper function to check if a subboard is valid
	check := func(board []int) bool {

		// map from a diagnoal direction to whether it's occupied
		// treating a row as a degenerate diagonal, with direction 0
		// a diagonal's id will be the row it crosses on the first column
		used := map[int]map[int]bool{-1: {}, 0: {}, 1: {}}

		// for every column, look at which rank the queen is in
		for col, rank := range board {
			// for each diagonal direction
			for dir, occupied := range used {
				// get the id of the diagonal
				diagonalID := rank - dir*col
				// if the diagonal is occupied, this board is invalid
				if occupied[diagonalID] {
					return false
					// otherwise, mark the diagonal as occupied
				} else {
					occupied[diagonalID] = true
				}
			}
		}

		// if we didn't collide on any diagonal this board is valid
		return true
	}

	// keep a record of the current board
	board := make([]int, n)

	// while we're looking at a column on the board
	for col := 0; col >= 0; {

		// move the queen forward by 1, resetting to 0 if we're past the end
		board[col] = (board[col] + 1) % (n + 1)

		// if the value is zero we've tested all options for this column
		// move to the prior column
		if board[col] == 0 {
			col--
			continue
		}

		// check if the columns we've set so far are valid
		valid := check(board[:col+1])

		// if they are valid and cover the full board, record an answer
		if valid && col == n-1 {
			ans = append(ans, append([]int{}, board...))
		} else if valid {
			// if it's valid but not complete, advance to next column
			col++
		}

	}

	return ans
}
