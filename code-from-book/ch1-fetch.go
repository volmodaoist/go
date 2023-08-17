package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
(1) 切到 go.work 文件所在目录，使用下面的两个编译命令
	不难看出 go build 其实能像 gcc 那样针对单个文件进行编译
		1.1. go build -o fetch ./code-from-book/ch1-fetch.go
		1.2. go build -o findlink ./code-from-book/ch5-findlinks1.go

(2) 终端运行 ./fetch http://gopl.io | ./findlink 即可看到运行的结果
	这个命令之中的 | 符号其实就是管道操作，其将 fetch 执行的结果作为输入丢给 findlink

(3) 如果想要不做编译，直接解释执行单个文件可以使用下列命令，或用管道来将前者的输出作为后者输入
	go run ./code-from-book/ch1-fetch.go 		http://gopl.io
	go run ./code-from-book/ch5-findlinks1.go
*/
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
