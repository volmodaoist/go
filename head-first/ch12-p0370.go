package main

import "fmt"

/**
 * 使用 recover 能够恢复panic，但是panic会阻断后续代码执行，因而recover必须
 * 配合 defer 使用才能发挥效果，不过 GO 设计者认为不要试图使用 panic和recover
 * 模拟try-catch-exception这种特性, 而是建议使用 if-else 以及函数err返回值处理
 */
func calmDown() {
	recover()
}

func freakout() {
	defer calmDown()
	panic("Holy shit!")
}

func main() {
	freakout()
	fmt.Println("Existing normally...")
}
