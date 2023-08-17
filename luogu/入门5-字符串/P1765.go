package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	var ans int
	nums := [...]int{
		1, 2, 3, 1, 2, 3, 1, 2, 3, 1,
		2, 3, 1, 2, 3, 1, 2, 3, 4, 1,
		2, 3, 1, 2, 3, 4}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := []rune(scanner.Text())

	for i := 0; i < len(str); i++ {
		if unicode.IsSpace(str[i]) {
			ans += 1
			continue
		}
		ans += nums[str[i]-'a']
	}

	fmt.Println(ans)
}
