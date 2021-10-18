package problems

import (
	"strconv"
	"strings"
)

type modStack struct {
	vals []int
	mins []int
}

func (ms *modStack) pop() int {
	v := ms.vals[len(ms.vals)-1]
	ms.vals = ms.vals[:len(ms.vals)-1]
	if v == ms.mins[len(ms.mins)-1] {
		ms.mins = ms.mins[:len(ms.mins)-1]
	}
	return v // don't think this matter for problem
}

func (ms *modStack) push(v int) {
	ms.vals = append(ms.vals, v)
	if len(ms.mins) == 0 || ms.mins[len(ms.mins)-1] >= v {
		ms.mins = append(ms.mins, v)
	}
}

func (ms *modStack) min() int {
	return ms.mins[len(ms.mins)-1]
}

func minimumOnStack(operations []string) []int {
	stack := modStack{[]int{}, []int{}}
	ans := []int{}
	for _, op := range operations {
		switch op := strings.Split(op, " "); op[0] {
		case "pop":
			stack.pop()
		case "push":
			v, _ := strconv.Atoi(op[1])
			stack.push(v)
		case "min":
			ans = append(ans, stack.min())
		}
	}
	return ans
}
