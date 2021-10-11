package problems

func findProfession(level int, pos int) string {

	pos -= 1
	flips := 0

	for pos > 0 {
		if pos&1 == 1 {
			flips++
		}
		pos >>= 1
	}

	if flips%2 == 0 {
		return "Engineer"
	} else {
		return "Doctor"
	}

}
