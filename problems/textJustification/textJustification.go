package textJustification

import (
	"strings"
)

var TextJustification = solution

func solution(words []string, l int) []string {

	// create an empty slice for the answer
	ans := []string{}

	// use i to mark which words should be part of a line
	// use j to mark which words are now part of a line
	for i, j := 0, 0; j < len(words); {

		// figure out how many words fit on this line
		wordsInLine, capacity := 0, l
		for i < len(words) && len(words[i]) <= capacity-wordsInLine {

			// add an additional word
			wordsInLine++

			// knock of the length of the word
			capacity -= len(words[i])

			// check the next word
			i++

		}

		// figure out if we're on the last line
		lastLine := i == len(words)

		// assume the spaces are size 1 with no extra
		spaceSize, excess := 1, 0

		// if there are multiple words and we're not on the last line
		// adjust the sizes to get full justification
		if wordsInLine > 1 && !lastLine {

			spacesInLine := wordsInLine - 1

			// quotient is the size of a "small" space
			spaceSize = capacity / spacesInLine

			// remainder is how many spaces need to be bigger by 1
			excess = capacity % spacesInLine
		}

		// build the line
		line := ""

		// while we've not added all the words we should to the line
		for ; j < i; j++ {

			// add the word to the line
			line += words[j]

			// if we're processing the last word, skip the spaces
			if j == i-1 {
				continue
			}

			// add in a "small" space
			line += strings.Repeat(" ", spaceSize)

			// if we haven't put in enough "big" spaces yet, make this one big
			if excess > 0 {
				line += " "
				excess--
			}

		}

		// pad out the rest of the line in case we didn't full justify it
		line += strings.Repeat(" ", l-len(line))

		// add this line to the answer
		ans = append(ans, line)

	}

	return ans

}
