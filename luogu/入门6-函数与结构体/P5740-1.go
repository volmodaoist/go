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

type SortByScoreIndex []Student

func (s SortByScoreIndex) Len() int {
	return len(s)
}

func (s SortByScoreIndex) Less(i, j int) bool {
	if s[i].score != s[j].score {
		return s[i].score > s[j].score
	} else {
		return s[i].index < s[j].index
	}
}

func (s SortByScoreIndex) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


/*
 *  非常不同于 C/C++一点在于 string 读入的时候也要使用&传递地址
 *  不仅如此，Scanf 在读入字符串的时候，有可能会读入换行符，若是混用 Scan 与 Scanf 极有可能出错！
*/
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
	sort.Sort(SortByScoreIndex(stus))
	fmt.Printf("%s %d %d %d\n", stus[0].name, stus[0].a, stus[0].b, stus[0].c)
}
