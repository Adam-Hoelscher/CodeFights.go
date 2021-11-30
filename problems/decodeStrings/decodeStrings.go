package decodeStrings

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {

	// start with a stack with one string
	stack := []string{""}

	// track whether we're in text mode
	text := true

	for _, c := range s {

		// if we shouldn't be in text mode, switch out
		if text && unicode.IsDigit(c) {
			text = false
			// make space for the coming numbers
			stack = append(stack, "")
		}

		// if the braces open, we're switching back to text mode
		if c == '[' {
			stack = append(stack, "")
			text = true

			// if the braces close, we need to do a multiplication
		} else if c == ']' {

			// pop the string to multiply and the multiplicand
			s := stack[len(stack)-1]
			n, _ := strconv.Atoi(stack[len(stack)-2])
			stack = stack[:len(stack)-2]

			// tack the multiplied string onto the tail of the stack
			stack[len(stack)-1] += strings.Repeat(s, n)

			// if we didn't do something with braces, grow the top item
		} else {
			stack[len(stack)-1] += string(c)

		}

	}

	return stack[0]
}
