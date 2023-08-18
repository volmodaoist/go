package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// 来自教材上面的案例，本节是对 Get 方法的一个封装
func main() {
	var url string
	fmt.Scan(&url)

	// 如果去掉下面这段代码，URL 有可能会因为缺少协议方案，e.g. "http://" 或 "https://"，陷入报错
	// WaitForServer 会因为 unsupported protocol scheme 错误处在等待状态
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v", err)
		os.Exit(1)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tris := 0; time.Now().Before(deadline); tris++ {
		doc, err := http.Get(url)
		if err == nil {
			fmt.Printf("Response type: %T\nResponse Value: %[1]v\n", doc)
			return nil
		}
		log.Printf("Serve not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tris))
	}
	return fmt.Errorf("Server %s failed to respond after %s", url, timeout)
}
