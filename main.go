package main

import (
	"fmt"
)

func main() {
	q := []int{1, 2, 3}
	for _, v := range q {
		fmt.Println(v)
		if v == 2 {
			q = append(q, 4)
		}
	}
	fmt.Println(q)
}
