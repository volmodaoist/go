package main

import (
	"fmt"
	"sort"
)

type Student struct {
	index          int
	name           string
	a, b, c, score int
}

func main() {
	var n int
	var stus []Student
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var name string
		var a, b, c int

		fmt.Scan(&name, &a, &b, &c)
		stus = append(stus, Student{i, name, a, b, c, a + b + c})
	}
	sort.Slice(stus, func(i, j int) bool {
		if stus[i].score != stus[j].score {
			return stus[i].score > stus[j].score
		} else {
			return stus[i].index < stus[j].index
		}
	})
	fmt.Printf("%s %d %d %d\n", stus[0].name, stus[0].a, stus[0].b, stus[0].c)
}
