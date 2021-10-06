package problems

func sudoku2(grid [][]string) bool {

	check := func(chars []string) bool {
		counter := map[string]int{}
		for _, char := range chars {
			counter[char] += 1
			if counter[char] > 1 {
				return false
			}
		}
		return true
	}

	grps := [3][9][]string{}

	for rowNum := 0; rowNum < 9; rowNum++ {
		for colNum := 0; colNum < 9; colNum++ {
			val := grid[rowNum][colNum]
			if val == "." {
				continue
			}
			sqrNum := 3*(rowNum/3) + (colNum / 3)
			grps[0][rowNum] = append(grps[0][rowNum], val)
			grps[1][colNum] = append(grps[1][colNum], val)
			grps[2][sqrNum] = append(grps[2][sqrNum], val)
		}
	}

	for _, grpSet := range grps {
		for _, grp := range grpSet {
			if !check(grp) {
				return false
			}
		}
	}

	return true
}
