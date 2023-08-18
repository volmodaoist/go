package main

import (
	"fmt"
)

type StuScore struct {
	name    string
	Chinese int
	Math    int
	English int
}

func sumScore(s StuScore) int {
	return s.Chinese + s.Math + s.English
}

func main() {
	var n, maxv int
	fmt.Scan(&n)
	list := make([]StuScore, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&list[i].name, &list[i].Chinese, &list[i].Math, &list[i].English)
		maxv = max(maxv, sumScore(list[i]))
	}

	for _, score := range list {
		if sumScore(score) == maxv {
			fmt.Printf("%s %d %d %d\n", score.name, score.Chinese, score.Math, score.English)
			break
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
