package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("c1 = %x\nc2 = %x\n", c1, c2)
	if c1 != c2 {
		fmt.Println("They are different!")
	}

	// 我们无法把一个数字转为一个 string 类型
	// 不过我们可以借助 Sprinf 函数得到任意类型的字符串格式
	s1 := fmt.Sprintf("%x", c1)
	s2 := fmt.Sprintf("%x", c2)
	fmt.Printf("Length of sha256-hash = %d\n", len(s1))
	count_different_bit(s1, s2)

	// 来自课本课后习题，测试的SHA256等等
	content := []byte("x")
	sha256Hash := ShaHash(content, 256)
	sha384Hash := ShaHash(content, 384)
	sha512Hash := ShaHash(content, 512)

	fmt.Printf("SHA-256 Hash: %x\n", sha256Hash)
	fmt.Printf("SHA-384 Hash: %x\n", sha384Hash)
	fmt.Printf("SHA-512 Hash: %x\n", sha512Hash)
}

// 来自课本的习题，统计两个 sha256-hash 散列值的不同位数
func count_different_bit(s1 string, s2 string) int {
	var cnt int
	for i := range s1 {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	fmt.Printf("One bit difference cause %d bits in hash!\n", cnt)
	return cnt
}

// 课后习题: 编写一个函数默认情况输出 SHA256, 但也支持 SHA384, SHA512
func ShaHash(content []byte, op int) []byte {
	if op == 256 {
		h := sha256.Sum256(content)
		return h[:]
	} else if op == 384 {
		h := sha512.Sum384(content)
		return h[:]
	} else if op == 512 {
		h := sha512.Sum512(content)
		return h[:]
	}
	// 返回空切片或其他错误处理
	return nil
}
